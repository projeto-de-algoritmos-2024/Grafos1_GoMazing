import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './MazeGeneration.css';

const TILE_SIZE = 50;

const MazeGeneration = () => {
    const [maze, setMaze] = useState([]);

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
        <div className="maze-container">
            {maze.map((row, rowIndex) =>
                row.map((cell, cellIndex) => (
                    <div
                        key={`${rowIndex}-${cellIndex}`}
                        className={`cell ${cell.visited ? 'visited' : ''}`}
                        style={{
                            top: rowIndex * TILE_SIZE,
                            left: cellIndex * TILE_SIZE,
                            width: TILE_SIZE,
                            height: TILE_SIZE,
                            borderTop: cell.Walls[0] ? '3px solid #1e4f5b' : 'none',
                            borderRight: cell.Walls[1] ? '3px solid #1e4f5b' : 'none',
                            borderBottom: cell.Walls[2] ? '3px solid #1e4f5b' : 'none',
                            borderLeft: cell.Walls[3] ? '3px solid #1e4f5b' : 'none',
                        }}
                    ></div>
                ))
            )}
        </div>
    );
};

export default MazeGeneration;