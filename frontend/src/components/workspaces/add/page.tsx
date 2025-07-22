// react imports
import { useState } from 'react';
import { MonsterAddElement, PDFMapMonster } from './monsterAddElement';
import { v4 as uuidGen } from 'uuid';

// wails imports
import { main } from '../../../../wailsjs/go/models';

export interface pdfMapPage {
    id: string,
    pageNumber: number,
    items: Array<PDFMapMonster>
}

interface PageNumberParameters {
    pageNumber: number,
    savePageNumber: (newPageNumber: number) => void
}
function PageNumber({pageNumber, savePageNumber} : PageNumberParameters) {
    const [pageNumberValue, setPageNumber] = useState(pageNumber);
    const [isEditingPageNumber, setEditingPageNumber] = useState(false);

    function pageNumberChange (event: React.ChangeEvent<HTMLInputElement>) {
        setPageNumber(Number(event.target.value));
    }

    function onClickSave() {
        setEditingPageNumber(false)
        savePageNumber(pageNumberValue)
    }

    if(isEditingPageNumber) {
        return (
            <div>
                <input onChange={pageNumberChange} value={pageNumberValue} type='number'></input>
                <button onClick={onClickSave}>save</button>
            </div>
        );
    }

    return <>
            {pageNumber}
            <div onClick={() => setEditingPageNumber(true)}>
                <i className='fa fa-edit'></i>
            </div>
        </>
}

export interface AddPageParameters{
    page: main.PDFMapPage,
    deletePage: (id: string) => void,
    saveChange: (newPage: main.PDFMapPage) => void
}
export function PageAddElement({page, saveChange, deletePage}: AddPageParameters) {


    function addNewItem() {
        let newItem: PDFMapMonster = {
            id: uuidGen(),
            name: '',
            cr: 0,
            monsterType: '',
            size: ''
        }
        let newItemList = page.items.map(item => item);
        newItemList.push(newItem);
        
        let newPage: main.PDFMapPage = new main.PDFMapPage;
        newPage.id= page.id;
        newPage.pageNumber = page.pageNumber;
        newPage.items = newItemList;

        saveChange(newPage);
    }

    function modifyItem(newItem: PDFMapMonster) {
        let newItems = page.items.map(item => {
            if(item.id === newItem.id) {
                return newItem
            }
            return item;
        });

        let newPage: main.PDFMapPage = new main.PDFMapPage;
        newPage.id= page.id;
        newPage.pageNumber = page.pageNumber;
        newPage.items = newItems;

        saveChange(newPage);
    }

    function deleteItem(id: string) {
        let newItems = page.items.filter(item => item.id !== id);
        let newPage: main.PDFMapPage = new main.PDFMapPage;
        newPage.id= page.id;
        newPage.pageNumber = page.pageNumber;
        newPage.items = newItems;

        saveChange(newPage);
    }

    function savePageNumber(newPageNumber: number) {
        let newPage: main.PDFMapPage = new main.PDFMapPage;
        newPage.id= page.id;
        newPage.pageNumber = newPageNumber;
        newPage.items = page.items;

        saveChange(newPage);
    }

    return (
       <div className='page'>
            <div className='page-header'>
                <PageNumber pageNumber={page.pageNumber} savePageNumber={savePageNumber}></PageNumber>
                <div className='spacer'></div>
                <div onClick={() => deletePage(page.id)}><i className='fa fa-trash'></i></div>
            </div>
            {page.items.map( entry => (
                <MonsterAddElement key={entry.id} monster={entry} saveItem={modifyItem} deleteItem={deleteItem}/>
            ))}
            <button onClick={() => addNewItem()}>Add Item</button>
       </div>
    );
}