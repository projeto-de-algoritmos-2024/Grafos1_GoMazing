import React, { useState, useEffect } from 'react';
import axios from 'axios';
import NodeComponent from './NodeComponent';
import './MazeApp.css';

const MazeGeneration = () => {
    const [maze, setMaze] = useState([]);
    const nodeSize = 50;

    useEffect(() => {
        generateMaze();
    }, []);

    const generateMaze = async () => {
        try {
            const response = await axios.get('http://localhost:8080/new-maze');
            setMaze(response.data);
        } catch (error) {
            console.error('Error generating maze:', error);
        }
    };

    return (
        <div className="maze-app">
            <h1>Maze Generation</h1>
            <div className="maze-grid" style={{ gridTemplateColumns: `repeat(${maze[0]?.length || 0}, ${nodeSize}px)` }}>
                {maze.map((row, rowIndex) =>
                    row.map((cell, cellIndex) => (
                        <NodeComponent
                            key={`${rowIndex}-${cellIndex}`}
                            nodeNumber={rowIndex * (maze[0]?.length || 0) + cellIndex}
                            size={nodeSize}
                            walls={cell.walls}
                            isActive={cell.visited}
                        />
                    ))
                )}
            </div>
        </div>
    );
};

export default MazeGeneration;