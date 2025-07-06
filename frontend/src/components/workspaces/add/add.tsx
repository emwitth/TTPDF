// css imports
import './add.scss';

// react imports
import {useState} from 'react';
import { MonsterAddElement, PDFMapMonster } from './monsterAddElement';
import { PageAddElement, pdfMapPage } from './page';

// go imports
import {GetPDF as GETPDF} from '../../../../wailsjs/go/main/App';

export function Add() {
    const [pdfURL, setPDFURL] = useState('');
    const [pdfName, setPDFName] = useState('');
    
    function GetPDF() {
        GETPDF().then(pdfURL => {
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
            <button className='button' onClick={GetPDF}>Choose PDF</button>
        </div>
    );
    
}

interface AddModifyWorkingAreaParameters {
    pdfName: string,
    pdfURL: string,
}
function AddModifyWorkingArea({pdfName, pdfURL} : AddModifyWorkingAreaParameters) {
    const [pdfMap, setPDFMap] = useState(new Array<pdfMapPage>);

    function addPDFMapEntry() {
        let pageNumber: number = pdfMap.length > 0 ? pdfMap[pdfMap.length - 1].pageNumber : 1;
        setPDFMap([
            ...pdfMap,
            {
                pageNumber: pageNumber,
                items: []
            }
        ]);
    }

    function addItem(newPage: pdfMapPage) {
        newPage.items.push(
            {
                name: "",
                cr: 0,
                type: "",
                size: "",
            }
        );
        let newPDFMap = pdfMap.map(page => {
            if(page.pageNumber === newPage.pageNumber) {
                return newPage;
            }
            return page
        })
        setPDFMap(newPDFMap);
    }
    
    return (
        <div>
            <div title={pdfURL}>
                Adding from {pdfName}
            </div>
            <div className='pages-container'>
                {pdfMap.map( page => (
                    <PageAddElement key={page.pageNumber} page={page} addItem={addItem}/>
                ))}
                <div className='page add-page' onClick={() => addPDFMapEntry()}>
                    <div style={{fontSize: "100px"}}>
                    +
                    </div>
                </div>
            </div>
        </div>
    );
}