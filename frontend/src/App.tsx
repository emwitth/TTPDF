// css imports
import './App.scss';

// react imports
import {createContext, useState} from 'react';
import { ActionBar } from './components/actionbar/actionBar';
import { Adventures } from './components/workspaces/adventures/adventures';
import { Monsters } from './components/workspaces/monsters/monsters';
import { Add } from './components/workspaces/add/add';

// go imports
import { GetSettings } from '../wailsjs/go/main/App';

enum WORKSPACES {
    Adventures,
    Monsters,
    Add
}

// default is just looking right here, like a normal web-app
export const PathContext = createContext('');

function App() {
    const [workspaceName, setWorkspaceName] = useState(WORKSPACES.Adventures);
    const [pathContext, setPathContext] = useState('');

    let workspaceHeight: string = document.body.scrollHeight + "px";
    GetSettings().then(settings => {
        setPathContext(settings.defaultPath);
    })

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
        <PathContext.Provider value={pathContext}>
            <div className='app'>
            <ActionBar onWorkspaceChange={onWorkspaceChange}/>
            <div className='workspace-area' style={{height: workspaceHeight}}>
                    <Workspace />
            </div>
            </div>
        </PathContext.Provider>
    )
}

export default App
