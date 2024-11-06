import React, { useEffect, useState } from 'react';
import './NodeComponent.css';

const NodeComponent = ({ nodeNumber, size, executionTime, walls }) => {
    const [style, setStyle] = useState({});
    const [text, setText] = useState('');

    useEffect(() => {
        setStyle({ width: `${size}px`, height: `${size}px` });
        setText(`Node number ${nodeNumber}.`);
    }, [nodeNumber, size]);

    return (
        <div className="node" style={style} title={text}>
            {nodeNumber}
            <div className={`wall top ${walls[0] ? 'active' : ''}`}></div>
            <div className={`wall right ${walls[1] ? 'active' : ''}`}></div>
            <div className={`wall bottom ${walls[2] ? 'active' : ''}`}></div>
            <div className={`wall left ${walls[3] ? 'active' : ''}`}></div>
        </div>
    );
};

export default NodeComponent;