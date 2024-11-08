import React, { useState, useEffect } from 'react';
import axios from 'axios';
import NodeComponent from './NodeComponent';
import './MazeApp.css';

const MazeApp = () => {
    const [width, setWidth] = useState(Math.trunc(window.innerWidth / 20));
    const [height, setHeight] = useState(Math.trunc(window.innerHeight / 20));
    const [maze, setMaze] = useState([]);
    const [steps, setSteps] = useState([]);
    const [currentStep, setCurrentStep] = useState(0);
    const [graph, setGraph] = useState([]);
    const nodeSize = 20;
    const executionTime = 10;

    useEffect(() => {
        populateGraph();
    }, [width, height]);

    useEffect(() => {
        const handleResize = () => {
            setWidth(Math.trunc(window.innerWidth / nodeSize));
            setHeight(Math.trunc(window.innerHeight / nodeSize));
        };

        window.addEventListener('resize', handleResize);
        return () => window.removeEventListener('resize', handleResize);
    }, []);

    const populateGraph = () => {
        const columns = Math.trunc(window.innerWidth / nodeSize);
        const rows = Math.trunc(window.innerHeight / nodeSize);
        let graph = [];
        let count = 0;

        for (let j = 0; j < rows; j++) {
            for (let i = 0; i < columns; i++) {
                graph[count] = [];
                if (j !== 0) {
                    graph[count].push(count - columns);
                }
                if (j !== rows - 1) {
                    graph[count].push(count + columns);
                }
                if (i !== 0) {
                    graph[count].push(count - 1);
                }
                if (i !== columns - 1) {
                    graph[count].push(count + 1);
                }
                count++;
            }
        }
        setGraph(graph);
    };

    const generateMaze = async (algo) => {
        try {
            const response = await axios.post('http://localhost:8080/generate', {
                width,
                height,
                algo
            });
            setSteps(response.data);
            setCurrentStep(0);
            setMaze(response.data[0]);
            animateMaze(response.data);
        } catch (error) {
            console.error('Error generating maze:', error);
        }
    };

    const animateMaze = (steps) => {
        let stepIndex = 0;
        const interval = setInterval(() => {
            if (stepIndex < steps.length) {
                setMaze(steps[stepIndex]);
                setCurrentStep(stepIndex); // Update current step
                stepIndex++;
            } else {
                clearInterval(interval);
            }
        }, 1500 / steps.length); // Adjust the speed as needed
    };

    const solveMaze = async (solveAlgo) => {
        try {
            const response = await axios.post('http://localhost:8080/solve', {
                algo: solveAlgo
            });
            setMaze(response.data);
        } catch (error) {
            console.error('Error solving maze:', error);
        }
    };

    return (
        <div className="maze-app">
            <h1>GoMazing</h1>
            <div className="controls">
                <label>
                    Width:
                    <input type="number" value={width} onChange={(e) => setWidth(e.target.value)} />
                </label>
                <label>
                    Height:
                    <input type="number" value={height} onChange={(e) => setHeight(e.target.value)} />
                </label>
            </div>
            <div className="buttons">
                <button onClick={() => generateMaze(1)}>Generate with DFS</button>
                <button onClick={() => generateMaze(2)}>Generate with Prim's</button>
                <button onClick={() => generateMaze(4)}>Generate with BFS</button>
            </div>
            {maze.length > 0 && (
                <div className="maze-display">
                    <h2>Maze</h2>
                    <div className="maze-grid" style={{ gridTemplateColumns: `repeat(${width}, ${nodeSize}px)` }}>
                        {maze.map((row, rowIndex) =>
                            row.map((cell, cellIndex) => (
                                <NodeComponent
                                    key={`${rowIndex}-${cellIndex}`}
                                    nodeNumber={rowIndex * width + cellIndex}
                                    size={nodeSize}
                                    walls={cell.walls}
                                    isActive={currentStep === rowIndex * width + cellIndex}
                                />
                            ))
                        )}
                    </div>
                    <button onClick={() => solveMaze(1)}>Solve with DFS</button>
                    <button onClick={() => solveMaze(2)}>Solve with BFS</button>
                </div>
            )}
        </div>
    );
};

export default MazeApp;