package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) getPageCount(path string) int {
	pageCount, err := api.PageCountFile(path)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "Error: failed to get file page count", err.Error())
		return 0
	}
	return pageCount
}

func (a *App) readPdfPage(path string, pageNumber int) io.Reader {
	// Open PDF file to parse
	ctx, err := api.ReadContextFile(path)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "Error: failed to read pdf", err.Error())
		return nil
	}

	page, err := api.ExtractPage(ctx, pageNumber)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "Error extracting PDF:", err.Error())
		return nil
	}

	return page
}

func (a *App) parsePdf(path string) error {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "pdfParsingDir")
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	defer os.RemoveAll(tempDir)

	// Extract pages into temporary directory
	err = api.ExtractPagesFile(path, tempDir, []string{"1-"}, nil)
	if err != nil {
		panic(err)
	}

	// Parse each PDF page
	name := strings.Split(filepath.Base(path), ".")
	totpages := a.getPageCount(path)
	for i := 1; i <= totpages; i++ {
		// <tempDirPath>/<filename>_page_<i>.pdf
		a.parsePdfPage(fmt.Sprintf("%s/%s_page_%d.pdf", tempDir, name[0], i), i)
	}

	return nil
}

func (a *App) parsePdfPage(path string, index int) error {
	fmt.Printf("Parsing PDF: %s\n", path)

	// Open (single page) PDF for text extraction
	page, reader, err := pdf.Open(path)
	if err != nil {
		panic(err)
	}
	defer page.Close()

	// Read plain text into string
	var buf bytes.Buffer
	ptreader, err := reader.GetPlainText()
	if err != nil {
		panic(err)
	}
	buf.ReadFrom(ptreader)
	text := buf.String()

	// Figure out what is going on with the text and handle it accordingly
	itemType := a.deducePdfMapItemType(text)
	switch itemType {
	case Monster5ePdfItemType:
		a.parse5eMonsterText(text)
	case UnknownPdfItemType:
		fallthrough
	default:
		fmt.Printf("Couldn't figure out the type of text from %s\n", path)
	}

	return nil
}
