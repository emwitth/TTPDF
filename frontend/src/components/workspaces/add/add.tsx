// css imports
import './add.scss';

// react imports
import {useState} from 'react';

// go imports
import {SelectPdf, GetPage, GetPageCount} from '../../../../wailsjs/go/main/App';

export function Add() {
    const [pdfData, setPdfData] = useState('');
    const [pdfUrl, setPdfUrl] = useState('');
    const [pdfName, setPdfName] = useState('');
    const [pageNumber, setPageNumber] = useState(-1);
    const [pageTotal, setPageTotal] = useState(-1);
    
    function ChoosePdf() {
        SelectPdf().then(url => {
            setPdfUrl(url);
            let name = JSON.stringify(url).split('\\').pop()?.split('.')[0];
            setPdfName(name ? name : '');
            UpdateTotalPages(url);
            LoadPage(url, 1);
        });
    }

    function UpdateTotalPages(url: string) {
        GetPageCount(url).then(totalPages => {
            setPageTotal(totalPages);
        })
    }

    function LoadPage(url: string, pageNum: number) {
        if (pageNum <= 0 || pageNum > pageTotal) {
            return;
        }
        GetPage(url, pageNum).then(data => {
            setPdfData("data:application/pdf;base64," + data + "#toolbar=0&navpanes=0");
            setPageNumber(pageNum);
        })
    }

    function AddButtonComponent({ChoosePDF} : {ChoosePDF: () => void}) {
        if (pdfData !== '') {
            return null;
        }

        return (
            <div>
                <div>Click to add a pdf</div><br />
                <button className='button' onClick={ChoosePDF}>Add PDF</button>
            </div>
        );
            
    }

    function PdfComponent() {
        if (pdfData === '') {
            return null;
        }

        return (
            <>
                <div>
                    <button className='button' onClick={() => {LoadPage(pdfUrl, pageNumber - 1)}}>
                        <i className='fa fa-minus'></i>
                    </button>
                    <div>{pageNumber} of {pageTotal}</div>
                    <button className='button' onClick={() => {LoadPage(pdfUrl, pageNumber + 1)}}>
                        <i className='fa fa-plus'></i>
                    </button>
                </div>
                <object data={pdfData} type="application/pdf" className='pdf'></object>
            </>
        )
    }

    return (
        <div className='add-container'>
            <div className='work-container'></div>
            <div className='pdf-container'>
                <AddButtonComponent ChoosePDF={ChoosePdf}/>
                <PdfComponent/>
            </div>
        </div>
    );
}