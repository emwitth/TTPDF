export interface PDFMapMonster
{
    name: string,
    cr: number
    type: string,
    size: string,

}

export interface AddMonsterElementParameters{

}
export function MonsterAddElement() {
    return (
        <div className="item">
            <table>
                <tbody>
                    <td>Name</td>
                    <td><input></input></td>
                </tbody>
                <tbody>
                    <td>CR</td>
                    <td><input></input></td>
                </tbody>
                <tbody>
                    <td>Type</td>
                    <td><input></input></td>
                </tbody>
                <tbody>
                    <td>Size</td>
                    <td><input></input></td>
                </tbody>
            </table>
        </div>
    );
}