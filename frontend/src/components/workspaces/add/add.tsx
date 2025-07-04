// css imports
import './add.scss';

// react imports
import {useState} from 'react';

// go imports
import {GetPDF as GETPDF} from "../../../../wailsjs/go/main/App";



export function Add() {
    const [pdf, setPDPF] = useState("");
    
    function GetPDF() {
        GETPDF().then(pdf => {
            setPDPF(pdf);
        })
    }
    
    return (
        <div className="add-div">
            <button className="button" onClick={GetPDF}>Add PDF</button>
            <div>{pdf}</div>
        </div>
    )
}