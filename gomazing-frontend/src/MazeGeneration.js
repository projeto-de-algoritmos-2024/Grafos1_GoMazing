import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './MazeGeneration.css';

const TILE_SIZE = 50;

const MazeGeneration = () => {
    const [mazeSteps, setMazeSteps] = useState([]);
    const [currentStep, setCurrentStep] = useState(0);
    const [algorithm, setAlgorithm] = useState(1); // Default to DFS

    useEffect(() => {
        generateMaze();
    }, [algorithm]);

    useEffect(() => {
        if (mazeSteps.length > 0) {
            const interval = setInterval(() => {
                setCurrentStep((prevStep) => {
                    if (prevStep < mazeSteps.length - 1) {
                        return prevStep + 1;
                    } else {
                        clearInterval(interval);
                        return prevStep;
                    }
                });
            }, 15000 / mazeSteps.length);
        }
    }, [mazeSteps]);

    const generateMaze = async () => {
        try {
            const response = await axios.post('http://localhost:8080/generate', {
                width: 24,
                height: 18,
                algo: algorithm
            });
            console.log("Maze generation response:", response.data); // Debugging line
            setMazeSteps(response.data);
        } catch (error) {
            console.error('Error generating maze:', error);
        }
    };

    return (
        <div className="maze-container">
            <div className="controls">
                <button onClick={() => setAlgorithm(1)}>Generate with DFS</button>
                <button onClick={() => setAlgorithm(2)}>Generate with BFS</button>
                <button onClick={() => setAlgorithm(3)}>Generate with Prim's</button>
            </div>
            {mazeSteps.length > 0 && mazeSteps[currentStep].map((row, rowIndex) =>
                Array.isArray(row) ? row.map((cell, cellIndex) => (
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
                )) : null
            )}
        </div>
    );
};

export default MazeGeneration;