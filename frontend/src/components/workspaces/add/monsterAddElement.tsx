// react imports
import { SyntheticEvent, useState } from 'react';

export interface PDFMapMonster
{
    id: string,
    name: string,
    cr: number
    monsterType: string,
    size: string,

}

export interface AddMonsterElementParameters{
    monster: PDFMapMonster,
    saveItem: (item: PDFMapMonster) => void,
    deleteItem: (id: string) => void
}
export function MonsterAddElement({monster, saveItem, deleteItem} : AddMonsterElementParameters) {
    const [isEditing, setEditing] = useState(false);
    const [name, setName] = useState(monster.name);
    const [cr, setCR] = useState(monster.cr);
    const [type, setType] = useState(monster.monsterType);
    const [size, setSize] = useState(monster.size);

    function nameChange (event: React.ChangeEvent<HTMLInputElement>) {
        setName(event.target.value);
    }

    function crChange (event: React.ChangeEvent<HTMLInputElement>) {
        setCR(Number(event.target.value));
    }

    function typeChange (event: React.ChangeEvent<HTMLInputElement>) {
        setType(event.target.value);
    }

    function sizeChange (event: React.ChangeEvent<HTMLInputElement>) {
        setSize(event.target.value);
    }

    function saveChanges() {
        let newMonsterItem = {
            id: monster.id,
            name: name,
            cr: cr,
            monsterType: type,
            size: size
        }

        saveItem(newMonsterItem);
        setEditing(false);
    }

    if(isEditing) {
        return (
            <div className='item'>
                <div className='item-table-area'>
                    <table>
                        <tbody>
                            <tr>
                                <td>Name</td>
                                <td><input onChange={nameChange} value={name}></input></td>
                            </tr>
                            <tr>
                                <td>CR</td>
                                <td><input onChange={crChange} value={cr} type='number'></input></td>
                            </tr>
                            <tr>
                                <td>Type</td>
                                <td><input onChange={typeChange} value={type}></input></td>
                            </tr>
                            <tr>
                                <td>Size</td>
                                <td><input onChange={sizeChange} value={size}></input></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div className='item-button-area'>
                    <button className='item-button-grow'  onClick={() => saveChanges()}>Save</button>
                    <button className='item-button-grow' onClick={() => deleteItem(monster.id)}>Delete</button>
                </div>
            </div>
        )
    }
    

    return (
        <div className='item'>
            <div className='item-table-area'>
                <table>
                    <tbody>
                        <tr>
                            <td>Name</td>
                            <td>{name}</td>
                        </tr>
                        <tr>
                            <td>CR</td>
                            <td>{cr}</td>
                        </tr>
                        <tr>
                            <td>Type</td>
                            <td>{type}</td>
                        </tr>
                        <tr>
                            <td>Size</td>
                            <td>{size}</td>
                        </tr>
                    </tbody>
                </table>
            </div>
            <div className='item-button-area'>
                <button className='item-button-grow' onClick={() => setEditing(true)}>Edit</button>
            </div>
        </div>
    );
}