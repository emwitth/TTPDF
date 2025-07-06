// css imports
import './App.scss';

// react imports
import {useState} from 'react';
import { ActionBar } from './components/actionbar/actionBar';
import { Adventures } from './components/workspaces/adventures/adventures';
import { Monsters } from './components/workspaces/monsters/monsters';
import { Add } from './components/workspaces/add/add';

enum WORKSPACES {
    Adventures,
    Monsters,
    Add
}

function App() {
    const [workspaceName, setWorkspaceName] = useState(WORKSPACES.Adventures);

    let workspaceHeight: string = document.body.scrollHeight + "px";


    function onWorkspaceChange(name: string) {
        switch(name) {
            case "Adventures":
                setWorkspaceName(WORKSPACES.Adventures);
                break;
            case "Monsters":
                setWorkspaceName(WORKSPACES.Monsters);
                break;
            case "Add":
                setWorkspaceName(WORKSPACES.Add);
                break;
        }
    }

    function Workspace() {
        switch(workspaceName) {
            case WORKSPACES.Adventures:
                return <Adventures/>
            case WORKSPACES.Monsters:
                return <Monsters/>
            case WORKSPACES.Add:
                return <Add/>
        }
    }

    return (
        <div className='app'>
           <ActionBar onWorkspaceChange={onWorkspaceChange}/>
           <div className='workspace-area' style={{height: workspaceHeight}}>
                <Workspace />
           </div>
        </div>
    )
}

export default App
