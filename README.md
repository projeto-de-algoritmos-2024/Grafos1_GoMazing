# GoMazing

**Número da Lista**: 57<br>
**Conteúdo da Disciplina**: Grafos 1<br>

## Alunos

| Matrícula   | Aluno                       |
| ----------- | --------------------------- |
| 19/0139323  | Pedro Menezes Rodiguero     |
| xx/xxxxxx   | xxxx xxxx xxxxx             |

## Sobre

GoMazing is a project that demonstrates various maze generation algorithms and a flood fill algorithm. The project includes a web interface where users can generate mazes using different algorithms and visualize the flood fill algorithm in action.

## Screenshots

![Maze Generation](screenshots/maze_generation.png)
![Flood Fill](screenshots/flood_fill.png)

## Instalação

**Linguagem**: Go, JavaScript (React)<br>
**Framework**: React, Gorilla Mux

### Pré-requisitos

- [Go](https://golang.org/doc/install) (versão 1.21.0 ou superior)
- [Node.js](https://nodejs.org/) (versão 14.0.0 ou superior)
- [npm](https://www.npmjs.com/get-npm) (versão 6.0.0 ou superior)

### Passos para instalação

1. Clone o repositório:

```sh
git clone https://github.com/projeto-de-algoritmos-2024/Grafos1_GoMazing.git
cd Grafos1_GoMazing
```
2. Instale as dependências do backend:
```sh
go mod tidy
```
3. Instale as dependências do frontend:
```sh
cd gomazing-frontend
npm install
```

### Uso

#### Executando o projeto
1. Inicie a aplicação
```sh
./run.sh
```

### Funcionalidades
- Maze Generation: Gere labirintos usando diferentes algoritmos (DFS, Prim's, Kruskal's, BFS).
- Flood Fill: Visualize o algoritmo de preenchimento de área em ação.

### Estrutura do projeto
```sh
.gitignore
algorithms/
    [bfs.go](http://_vscodecontentref_/2)
    [dfs.go](http://_vscodecontentref_/3)
    [helpers.go](http://_vscodecontentref_/4)
    [kruskal.go](http://_vscodecontentref_/5)
    [prim.go](http://_vscodecontentref_/6)
    [solve.go](http://_vscodecontentref_/7)
[go.mod](http://_vscodecontentref_/8)
[go.sum](http://_vscodecontentref_/9)
gomazing-frontend/
    .gitignore
    [package.json](http://_vscodecontentref_/10)
    public/
        [index.html](http://_vscodecontentref_/11)
        [manifest.json](http://_vscodecontentref_/12)
        [robots.txt](http://_vscodecontentref_/13)
    [README.md](http://_vscodecontentref_/14)
    src/
        [App.css](http://_vscodecontentref_/15)
        [App.js](http://_vscodecontentref_/16)
        [App.test.js](http://_vscodecontentref_/17)
        [FloodFillApp.css](http://_vscodecontentref_/18)
        [FloodFillApp.js](http://_vscodecontentref_/19)
        [index.css](http://_vscodecontentref_/20)
        [index.js](http://_vscodecontentref_/21)
        [MazeApp.css](http://_vscodecontentref_/22)
        [MazeApp.js](http://_vscodecontentref_/23)
        [NodeComponent.css](http://_vscodecontentref_/24)
        [NodeComponent.js](http://_vscodecontentref_/25)
        [reportWebVitals.js](http://_vscodecontentref_/26)
        [setupTests.js](http://_vscodecontentref_/27)
[README.md](http://_vscodecontentref_/28)
[run.sh](http://_vscodecontentref_/29)
[server.go](http://_vscodecontentref_/30)
```

### Outros
Quaisquer outras informações sobre seu projeto podem ser descritas abaixo.

- Licença: MIT
