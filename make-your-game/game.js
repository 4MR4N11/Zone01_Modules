const board = document.getElementById('game-board')

const ROWS = 20
const COLUMNS = 10

let grid = Array.from({ length: ROWS }, () => Array(COLUMNS).fill(0))

const tetrominoes = [
    [[1, 1, 1], [0, 1, 0]],  // T-shape
    [[1, 1], [1, 1]],        // O-shape
    [[1, 0, 0], [1, 1, 1]],  // L-shape
    [[0, 0, 1], [1, 1, 1]],  // J-shape
    [[0, 1, 1], [1, 1, 0]],  // S-shape
    [[1, 1, 0], [0, 1, 1]],  // Z-shape
    [[1, 1, 1, 1]],          // I-shape
]

let currentPiece = tetrominoes[0]
let x = 0
let y = 4
let dropStart = 0 
let dropInterval = 400 //ms
let pause = false
let gameover = false
let score = 0
let timeElapsed = 0;
let lastTimerUpdate = 0;

function renderGrid() {
    board.innerHTML = ''
    const scoreElement = document.createElement('div')
    const banner = document.createElement('div')
    banner.className = 'banner'
    scoreElement.className = 'score'
    scoreElement.textContent = `Score: ${score}`
    const timerElement = document.createElement('div')
    timerElement.id = 'timer'
    timerElement.textContent = `Time: ${timeElapsed}`
    banner.style.display = 'flex'
    banner.style.justifyContent = 'space-between'
    banner.appendChild(scoreElement)
    banner.appendChild(timerElement)
    board.appendChild(banner)

    for (let row = 0; row < ROWS; row++) {
        const rowElement = document.createElement('div')
        rowElement.className = 'row'
        for (let col = 0; col < COLUMNS; col++) {
            const cell = document.createElement('div')
            cell.className = 'cell'
            if (grid[row][col] === 1) {
                cell.classList.add('active')
            } else if (grid[row][col] === 2) {
                cell.classList.add('fixed')
            }
            rowElement.appendChild(cell)
        }
        board.appendChild(rowElement)
    }
}

function placeTetromino() {
    clearLines()
    if (checkGameOver()) {
        gameover = true
        return
    }
    for (let row = 0; row < currentPiece.length; row++) {
        for (let col = 0; col < currentPiece[row].length; col++) {
            if (currentPiece[row][col] === 1) {
                grid[x + row][y + col] = 1
            }
        }
    }
}

function removeTetromino() {
    for (let row = 0; row < currentPiece.length; row++) {
        for (let col = 0; col < currentPiece[row].length; col++) {
            if (currentPiece[row][col] === 1) {
                grid[x + row][y + col] = 0
            }
        }
    }
}

function clearLines() {
    const lines = document.querySelectorAll('.row')
    for (let line = 0; line < lines.length; line++) {
        const cells = lines[line].querySelectorAll('.cell')
        let isFull = true
        for (let cell = 0; cell < cells.length; cell++) {
            if (!cells[cell].classList.contains('fixed')) {
                isFull = false
                break
            }
        }
        if (isFull) {
            for (let row = line; row > 0; row--) {
                for (let col = 0; col < COLUMNS; col++) {
                    grid[row][col] = grid[row - 1][col]
                }
            }
            score++
        }
    }

}


function checkGameOver() {
    for (let col = 0; col < COLUMNS; col++) {
        if (grid[0][col] === 2) {
            return true
        }
    }
    return false
}

function moveDown() {
    removeTetromino()
    x++
    if (!isValidMove()) {
        x--
        fixTetromino()
        spawnTetromino()
    }
    placeTetromino()
    renderGrid()
}

function isValidMove() {
    for (let row = 0; row < currentPiece.length; row++) {
        for (let col = 0; col < currentPiece[row].length; col++) {
            if (
                currentPiece[row][col] === 1 &&
                (x + row >= ROWS ||
                    y + col < 0 ||
                    y + col >= COLUMNS ||
                    grid[x + row][y + col] === 2)
            ) {
                return false
            }
        }
    }
    return true
}

function fixTetromino() {
    for (let row = 0; row < currentPiece.length; row++) {
        for (let col = 0; col < currentPiece[row].length; col++) {
            if (currentPiece[row][col] === 1) {
                grid[x + row][y + col] = 2
            }
        }
    }
}

function spawnTetromino() {
    currentPiece = tetrominoes[Math.floor(Math.random() * tetrominoes.length)]
    x = 0
    y = Math.floor(COLUMNS / 2) - Math.floor(currentPiece[0].length / 2)
}

function rotatePiece(counterClockwise = false) {
    const rows = currentPiece.length
    const cols = currentPiece[0].length
    const rotated = Array.from({ length: cols }, () => Array(rows).fill(0))

    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            if (counterClockwise) {
                rotated[cols - col - 1][row] = currentPiece[row][col]
            } else {
                rotated[col][rows - row - 1] = currentPiece[row][col]
            }
        }
    }

    currentPiece = rotated
}

document.addEventListener('keydown', (event) => {
    switch (event.key) {
        case 'ArrowLeft':
            removeTetromino()
            y--
            if (!isValidMove()) y++
            placeTetromino()
            renderGrid()
            break

        case 'ArrowRight':
            removeTetromino()
            y++
            if (!isValidMove()) y--
            placeTetromino()
            renderGrid()
            break

        case 'ArrowDown':
            moveDown()
            break

        case 'ArrowUp':
            removeTetromino()
            rotatePiece()
            if (!isValidMove()) rotatePiece(true)
            placeTetromino()
            renderGrid()
            break

        case ' ':
            removeTetromino()
            while (isValidMove()) {
                x++
            }
            x--
            placeTetromino()
            fixTetromino()
            spawnTetromino()
            renderGrid()
            break
        case 'Escape':
            pauseMenu()
            pause = !pause
            if (!pause) {
                requestAnimationFrame(gameLoop)
            }
            break

    }
})

function pauseMenu() {
    const pauseMenu = document.createElement('div')
    pauseMenu.className = 'pause-menu'
    const pauseText = document.createElement('h1')
    pauseText.textContent = 'Paused'
    pauseMenu.appendChild(pauseText)
    board.appendChild(pauseMenu)
}


function updateTimer(timestamp) {
    if (timestamp - lastTimerUpdate >= 1000) {
        timeElapsed++;
        lastTimerUpdate = timestamp;
        document.getElementById('timer').textContent = `Time: ${timeElapsed}`;
    }
}

function gameLoop(timestamp) {
    updateTimer(timestamp)
    if (gameover) {
        alert('Game Over')
        return
    }
    if (pause) return
    if (!dropStart) dropStart = timestamp
    const time = timestamp - dropStart
    if (time > dropInterval) {
        moveDown()
        dropStart = timestamp
    }
    requestAnimationFrame(gameLoop)
}

placeTetromino()
renderGrid()
requestAnimationFrame(gameLoop)
