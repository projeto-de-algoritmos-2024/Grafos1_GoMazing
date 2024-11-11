import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';
import './MazeGeneration.css';

const TILE_SIZE = 20;
const COLORS = ['#1e1e1e', '#f70067', '#3eb489']; // Define colors for each pass

const MazeGeneration = () => {
    const [mazeSteps, setMazeSteps] = useState([]);
    const [currentStep, setCurrentStep] = useState(0);
    const [algorithm, setAlgorithm] = useState(null); // No default algorithm

    const generateMaze = useCallback(async () => {
        if (algorithm === null) return; // Ensure an algorithm is selected

        try {
            const response = await axios.post('http://localhost:8080/generate', {
                width: 60,
                height: 30,
                algo: algorithm
            });
            console.log("Maze generation response:", response.data); // Debugging line
            setMazeSteps(response.data);
        } catch (error) {
            console.error('Error generating maze:', error);
        }
    }, [algorithm]);

    useEffect(() => {
        if (mazeSteps.length > 0) {
            const intervalTime = algorithm === 2 ? 5000 / mazeSteps.length : 15000 / mazeSteps.length; // Faster for BFS
            const interval = setInterval(() => {
                setCurrentStep((prevStep) => {
                    if (prevStep < mazeSteps.length - 1) {
                        return prevStep + 1;
                    } else {
                        clearInterval(interval);
                        return prevStep;
                    }
                });
            }, intervalTime);
        }
    }, [mazeSteps, algorithm]);

    return (
        <div className="maze-container">
            <div className="controls">
                <button onClick={() => setAlgorithm(1)}>Select DFS</button>
                <button onClick={() => setAlgorithm(2)}>Select BFS</button>
                <button onClick={() => setAlgorithm(3)}>Select asdasdasd0's</button>
                <button onClick={generateMaze} disabled={algorithm === null}>Generate Maze</button>
            </div>
            {mazeSteps.length > 0 && mazeSteps[currentStep].map((row, rowIndex) =>
                Array.isArray(row) ? row.map((cell, cellIndex) => {
                    const colorIndex = Math.floor(currentStep / (mazeSteps.length / COLORS.length)) % COLORS.length;
                    return (
                        <div
                            key={`${rowIndex}-${cellIndex}`}
                            className={`cell ${cell.visited ? 'visited' : ''}`}
                            style={{
                                top: rowIndex * TILE_SIZE,
                                left: cellIndex * TILE_SIZE,
                                width: TILE_SIZE,
                                height: TILE_SIZE,
                                backgroundColor: cell.visited ? COLORS[colorIndex] : 'transparent',
                                borderTop: cell.Walls[0] ? '3px solid #1e4f5b' : 'none',
                                borderRight: cell.Walls[1] ? '3px solid #1e4f5b' : 'none',
                                borderBottom: cell.Walls[2] ? '3px solid #1e4f5b' : 'none',
                                borderLeft: cell.Walls[3] ? '3px solid #1e4f5b' : 'none',
                            }}
                        ></div>
                    );
                }) : null
            )}
        </div>
    );
};

export default MazeGeneration;