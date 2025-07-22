// css imports
import './add.scss';

// react imports
import {useState} from 'react';
import { PageAddElement, pdfMapPage } from './page';
import { v4 as uuidGen } from 'uuid';

// go imports
import {GetPDF as GETPDF} from '../../../../wailsjs/go/main/App';
import {SaveCSV, SavePDFMap, OpenPDFMap} from '../../../../wailsjs/go/main/App';
import { main } from '../../../../wailsjs/go/models';

export function Add() {
    const [pdfURL, setPDFURL] = useState('');
    const [pdfName, setPDFName] = useState('');
    
    function GetPDF() {
        GETPDF().then((pdfURL) => {
            setPDFURL(pdfURL);
            let pdfName = JSON.stringify(pdfURL).split('\\').pop()?.split('.')[0];
            setPDFName(pdfName ? pdfName : '');
        });
    }

     if(pdfName === '') {
        return (
        <div className='add-container'>
            <div className='work-container work-container-centered-vertical'>
                <AddButtonComponent GetPDF={GetPDF} pdfName={pdfName} />
            </div>
        </div>
    );
    }

    return (
        <div className='add-container'>
            <div className='work-container'>
                <AddModifyWorkingArea pdfName={pdfName} pdfURL={pdfURL}/>
            </div>
        </div>
    );
}

interface AddButtonComponentParameters {
    GetPDF: () => void,
    pdfName: string
}
function AddButtonComponent({GetPDF, pdfName} : AddButtonComponentParameters) {
    return (
        <div>
            <div>Click to create or modify a pdf mapping</div><br />
            <button onClick={GetPDF}>Choose PDF</button>
        </div>
    );
    
}

interface AddModifyWorkingAreaParameters {
    pdfName: string,
    pdfURL: string,
}
function AddModifyWorkingArea({pdfName, pdfURL} : AddModifyWorkingAreaParameters) {
    const [pdfMap, setPDFMap] = useState(new Array<main.PDFMapPage>);

    function addPage() {
        let pageNumber: number = pdfMap.length > 0 ? pdfMap[pdfMap.length - 1].pageNumber + 1 : 1;
        let newPDFMapPage = new main.PDFMapPage();
        newPDFMapPage.id = uuidGen();
        newPDFMapPage.pageNumber = pageNumber;
        newPDFMapPage.items = [];
        setPDFMap([
            ...pdfMap,
           newPDFMapPage
        ]);
    }

    function deletePage(id: string) {
        let newPDFMap = pdfMap.filter(page => page.id !== id);
        setPDFMap(newPDFMap);
    }

    function saveItem(newPage: main.PDFMapPage) {
        console.log(newPage);
        let newPDFMap = pdfMap.map(page => {
            if(page.id === newPage.id) {
                return newPage;
            }
            return page
        })
        console.log(newPDFMap);
        setPDFMap(newPDFMap);
        SavePDFMap(newPDFMap, pdfURL);
    }
    
    return (
        <div>
            <div title={pdfURL}>
                Adding from {pdfName}
            </div>
            <div className='pages-container'>
                {pdfMap.map( (page) => (
                    <PageAddElement key={page.id} page={page} saveChange={saveItem} deletePage={deletePage}/>
                ))}
                <div className='page add-page' onClick={() => addPage()}>
                    <div style={{fontSize: "100px"}}>
                        <i className='fa fa-plus'></i>
                    </div>
                </div>
            </div>
        </div>
    );
}