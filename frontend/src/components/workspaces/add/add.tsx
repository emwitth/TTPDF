// css imports
import './add.scss';

// react imports
import {useState, useContext} from 'react';
import { PathContext } from '../../../App';

// go imports
import {SelectPdf} from '../../../../wailsjs/go/main/App';

export function Add() {
    const [pdfUrl, setPdfUrl] = useState('');
    const [pdfRelativePath, setPdfRelativePath] = useState('');
    const [pdfName, setPdfName] = useState('');
    const pathContext = useContext(PathContext);
    
    function ChoosePdf() {
        SelectPdf().then(url => {
            setPdfUrl(url);
            let name = JSON.stringify(url).split('\\').pop()?.split('.')[0];
            setPdfName(name ? name : '');
            let relPath = JSON.stringify(url).split(pathContext).pop();
            // we need the substring because wails is stupid and decided to send back the path in quotes
            setPdfRelativePath(relPath ? relPath.substring(0, relPath.length-1) : '');
            console.log(relPath);
        });
    }

    function AddButtonComponent({ChoosePDF} : {ChoosePDF: () => void}) {
        if (pdfRelativePath !== '') {
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
        if (pdfRelativePath === '') {
            return null;
        }

        return (
            <>
                <div>
                </div>
                <object data={pdfRelativePath} type="application/pdf" className='pdf'></object>
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