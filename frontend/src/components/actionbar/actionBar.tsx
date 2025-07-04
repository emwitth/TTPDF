// css imports
import './actionbar.scss';

// react imports
import { useState } from 'react';

interface ActionBarData {
    onWorkspaceChange: (name: string) => void
}

interface ActionBarButtonData {
    name : string,
    isActive : boolean,
    onButtonClick: React.MouseEventHandler<HTMLDivElement>
}

let initialActionState = [
    {
        name: "Adventures",
        isActive: true
    },
    {
        name: "Monsters",
        isActive: false
    },
    {
        name: "Add",
        isActive: false
    },
]

export function ActionBar({onWorkspaceChange}: ActionBarData) {
    const [actions, setActions] = useState(initialActionState);

    function handleClick(actionName: string) {
        if(!actions.find((action) => { action.name === actionName})?.isActive) {
            let newActions = actions.map((action) => {
                let isActive = false;
                if (action.name === actionName) {
                    isActive = true;
                }
                return {
                    name: action.name,
                    isActive: isActive
                };
            });
            setActions(newActions);
            onWorkspaceChange(actionName);
        }
    }

    return (
        <div className="action-bar">
            {actions.map(action => (
                <ActionBarButton key={action.name} name={action.name} isActive={action.isActive} onButtonClick={() => handleClick(action.name)}/>
            ))}
        </div>
    );
}

function ActionBarButton({name, isActive, onButtonClick} : ActionBarButtonData) {
    let className: string = "action-bar-button";
    className += isActive ? " active" : " inactive"

    return <div className={className} onClick={onButtonClick}>{name}</div>
}