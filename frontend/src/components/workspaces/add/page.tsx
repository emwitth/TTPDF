// react imports
import { MonsterAddElement, PDFMapMonster } from "./monsterAddElement";

export interface pdfMapPage {
    pageNumber: number,
    items: Array<PDFMapMonster>
}

export interface AddPageParameters{
    page: pdfMapPage,
    addItem: (page: pdfMapPage) => void
}
export function PageAddElement({page, addItem}: AddPageParameters) {
    return (
       <div className="page">
        {page.items.map( entry => (
            <MonsterAddElement key={entry.name} />
        ))}
            <button className='button' onClick={() => addItem(page)}>Add Item</button>
       </div>
    );
}