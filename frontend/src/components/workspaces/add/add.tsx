// css imports
import './add.scss';

// react imports
import {useState} from 'react';

// go imports
import {GetPDF as GETPDF} from '../../../../wailsjs/go/main/App';

export function Add() {
    const [pdfURL, setPDFURL] = useState('');
    const [pdf, setPDF] = useState('');
    
    function GetPDF() {
        GETPDF().then(pdfURL => {
            setPDFURL(pdfURL);
            let pdf = JSON.stringify(pdfURL).split('\\').pop()?.split('.')[0];
            setPDF(pdf ? pdf : '');
        });
    }

    function AddButtonComponent({GetPDF} : {GetPDF: () => void}) {
        if (pdf !== '') {
            return null;
        }

        return (
            <div>
                <div>Click to add a pdf</div><br />
                <button className='button' onClick={GetPDF}>Add PDF</button>
            </div>
        );
            
    }

    function AddMonsterFormComponent({}) {
        if(pdf === '') {
            return null;
        }

        return (
            <div>
                <div title={pdfURL}>
                    Adding from {pdf};
                </div>
                CR <input></input>
            </div>
        );
    }

    return (
        <div className='add-container'>
            <div className='work-container'>
                <AddButtonComponent GetPDF={GetPDF} />
                <AddMonsterFormComponent/>
            </div>
        </div>
    );
}