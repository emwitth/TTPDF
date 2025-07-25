package main

import (
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) getPageCount(filePath string) int {
	pageCount, err := api.PageCountFile(filePath)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "Error: failed to get file page count", err.Error())
		return 0
	}
	return pageCount
}

func (a *App) readPdfPage(filePath string, pageNumber int) io.Reader {
	ctx, err := api.ReadContextFile(filePath)
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
