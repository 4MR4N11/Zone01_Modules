const board = document.getElementById("game-board");

const ROWS = 20;
const COLUMNS = 10;

let grid = Array.from({ length: ROWS }, () => Array(COLUMNS).fill(0));

const tetrominoes = [
  [
    [0, 1, 0],
    [1, 1, 1],
  ], // T-shape
  [
    [1, 1],
    [1, 1],
  ], // O-shape
  [
    [1, 0, 0],
    [1, 1, 1],
  ], // L-shape
  [
    [0, 0, 1],
    [1, 1, 1],
  ], // J-shape
  [
    [0, 1, 1],
    [1, 1, 0],
  ], // S-shape
  [
    [1, 1, 0],
    [0, 1, 1],
  ], // Z-shape
  [[1, 1, 1, 1]], // I-shape
];

const colors = [
  "yellow",
  "blue",
  "brown",
  "red",
  "chocolate",
  "purple",
  "cyan",
];

let index = Math.floor(Math.random() * tetrominoes.length);
let currentPiece = tetrominoes[index];
let color = colors[Math.floor(Math.random() * colors.length)];
let x = 0;
let y = 4;
let dropStart = 0;
let dropInterval = 400; //ms
let start = false;
let pause = false;
let gameover = false;
let score = 0;
let scoreInc = 13;
let timeElapsed = 0;
let lastTimerUpdate = 0;

function initGrid() {
  board.innerHTML = "";
  const scoreElement = document.createElement("div");
  const banner = document.createElement("div");
  banner.className = "banner";
  scoreElement.className = "score";
  scoreElement.textContent = `Score: ${score}`;
  const timerElement = document.createElement("div");
  timerElement.id = "timer";
  timerElement.textContent = `Time: ${timeElapsed}`;
  banner.style.display = "flex";
  banner.style.justifyContent = "space-between";
  banner.appendChild(scoreElement);
  banner.appendChild(timerElement);
  board.appendChild(banner);
  for (let row = 0; row < ROWS; row++) {
    const rowElement = document.createElement("div");
    rowElement.className = "row";
    for (let col = 0; col < COLUMNS; col++) {
      const cell = document.createElement("div");
      cell.className = "cell";
      if (grid[row][col] === 1) {
        cell.classList.add("active");
        cell.style.backgroundColor = color;
      } else if (grid[row][col] === 2) {
        cell.classList.add("fixed");
      }
      rowElement.appendChild(cell);
    }
    board.appendChild(rowElement);
  }
}

function renderGrid() {
  const rows = board.querySelectorAll(".row");
  for (let row = 0; row < rows.length; row++) {
    const rowElement = rows[row];
    const cells = rowElement.querySelectorAll(".cell");
    for (let col = 0; col < cells.length; col++) {
      const cell = cells[col];
      cell.className = "cell";
      cell.style = {};
      if (grid[row][col] === 1) {
        cell.classList.add("active");
        cell.style.backgroundColor = color;
      } else if (grid[row][col] === 2) {
        cell.classList.add("fixed");
      }
    }
  }
}

function placeTetromino() {
  clearLines();
  if (checkGameOver()) {
    gameover = true;
    return;
  }
  for (let row = 0; row < currentPiece.length; row++) {
    for (let col = 0; col < currentPiece[row].length; col++) {
      if (currentPiece[row][col] === 1) {
        grid[x + row][y + col] = 1;
      }
    }
  }
}

function removeTetromino() {
  for (let row = 0; row < currentPiece.length; row++) {
    for (let col = 0; col < currentPiece[row].length; col++) {
      if (currentPiece[row][col] === 1) {
        grid[x + row][y + col] = 0;
      }
    }
  }
}

function clearLines() {
  const lines = document.querySelectorAll(".row");
  for (let line = 0; line < lines.length; line++) {
    const cells = lines[line].querySelectorAll(".cell");
    let isFull = true;
    for (let cell = 0; cell < cells.length; cell++) {
      if (!cells[cell].classList.contains("fixed")) {
        isFull = false;
        break;
      }
    }
    if (isFull) {
      grid.splice(line, 1);
      grid.unshift(Array.from({ length: COLUMNS }).fill(0));
      score += scoreInc;
      document.querySelector(".score").textContent = `Score: ${score}`;
    }
  }
}

function checkGameOver() {
  for (let col = 0; col < COLUMNS; col++) {
    if (grid[0][col] === 2) {
      return true;
    }
  }
  return false;
}

function moveDown() {
  removeTetromino();
  x++;
  if (!isValidMove()) {
    x--;
    fixTetromino();
    spawnTetromino();
  }
  placeTetromino();
  renderGrid();
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
        return false;
      }
    }
  }
  return true;
}

function fixTetromino() {
  for (let row = 0; row < currentPiece.length; row++) {
    for (let col = 0; col < currentPiece[row].length; col++) {
      if (currentPiece[row][col] === 1) {
        grid[x + row][y + col] = 2;
      }
    }
  }
}

function spawnTetromino() {
  color = colors[Math.floor(Math.random() * tetrominoes.length)];
  index = Math.floor(Math.random() * tetrominoes.length);
  currentPiece = tetrominoes[index];
  x = 0;
  y = Math.floor(COLUMNS / 2) - Math.floor(currentPiece[0].length / 2);
}

function rotatePiece(counterClockwise = false) {
  const rows = currentPiece.length;
  const cols = currentPiece[0].length;
  const rotated = Array.from({ length: cols }, () => Array(rows).fill(0));
  for (let row = 0; row < rows; row++) {
    for (let col = 0; col < cols; col++) {
      if (counterClockwise) {
        rotated[cols - 1 - col][row] = currentPiece[row][col];
      } else {
        rotated[col][rows - 1 - row] = currentPiece[row][col];
      }
    }
  }
  currentPiece = rotated;
}

let gameState = false;
document.addEventListener("keydown", (event) => {
  switch (event.key) {
    case "ArrowLeft":
      if (gameState) {
        if (!pause) {
          removeTetromino();
          y--;
          if (!isValidMove()) y++;
          placeTetromino();
          renderGrid();
        }
      }
      break;
    case "ArrowRight":
      if (gameState) {
        if (!pause) {
          removeTetromino();
          y++;
          if (!isValidMove()) y--;
          placeTetromino();
          renderGrid();
        }
      }
      break;

    case "ArrowDown":
      if (gameState) {
        if (!pause) {
          moveDown();
        }
      }
      break;

    case "ArrowUp":
      if (gameState) {
        if (!pause) {
          removeTetromino();
          rotatePiece();
          if (!isValidMove()) rotatePiece(true);
          placeTetromino();
          renderGrid();
        }
      }
      break;

    case " ":
      if (gameState) {
        if (!pause) {
          removeTetromino();
          while (isValidMove()) {
            x++;
          }
          x--;
          placeTetromino();
          fixTetromino();
          spawnTetromino();
          renderGrid();
        }
      }
      break;
    case "Escape":
      if (gameState) {
        pauseMenu();
        pause = !pause;
        if (!pause) {
          requestAnimationFrame(gameLoop);
        }
      }
      break;
  }
});

function pauseMenu() {
  const pauseMenu = document.createElement("div");
  pauseMenu.className = "pause-menu";
  const pauseText = document.createElement("h1");
  pauseText.textContent = "Paused";
  const resumeButton = document.createElement("button");
  const restartButton = document.createElement("button");
  const buttons = document.createElement("div");
  restartButton.textContent = "Restart";
  restartButton.addEventListener("click", () => {
    location.reload();
  });
  resumeButton.textContent = "Resume";
  resumeButton.addEventListener("click", () => {
    pause = false;
    board.removeChild(pauseMenu);
    requestAnimationFrame(gameLoop);
  });
  buttons.appendChild(restartButton);
  buttons.appendChild(resumeButton);
  pauseMenu.appendChild(pauseText);
  pauseMenu.appendChild(buttons);
  board.appendChild(pauseMenu);
}

function updateTimer(timestamp) {
  if (timestamp - lastTimerUpdate >= 1000) {
    timeElapsed++;
    lastTimerUpdate = timestamp;
    document.getElementById("timer").textContent = `Time: ${timeElapsed}`;
  }
}

function startMenu() {
  const startMenu = document.createElement("div");
  startMenu.className = "startMenu";
  const startText = document.createElement("h1");
  startText.textContent = "New Game";
  const startButton = document.createElement("button");
  const button = document.createElement("div");
  startButton.textContent = "Start";
  button.appendChild(startButton);
  startMenu.appendChild(startText);
  startMenu.appendChild(button);
  board.appendChild(startMenu);
  startButton.addEventListener("click", () => {
    gameState = true;
    board.removeChild(startMenu);
    requestAnimationFrame(gameLoop);
  });
}

function gameOver() {
  const startMenu = document.createElement("div");
  startMenu.className = "startMenu";
  const startText = document.createElement("h1");
  startText.textContent = "Game Over";
  const startButton = document.createElement("button");
  const button = document.createElement("div");
  startButton.textContent = "Restart?";
  button.appendChild(startButton);
  startMenu.appendChild(startText);
  startMenu.appendChild(button);
  board.appendChild(startMenu);
  startButton.addEventListener("click", () => {
    location.reload();
  });
}

function gameLoop(timestamp) {
  updateTimer(timestamp);
  if (gameover) {
    gameState = !gameState;
    gameOver();
    return;
  }
  if (pause) return;
  if (!dropStart) dropStart = timestamp;
  const time = timestamp - dropStart;
  if (time > dropInterval) {
    moveDown();
    dropStart = timestamp;
  }
  requestAnimationFrame(gameLoop);
}
initGrid();
placeTetromino();
renderGrid();
startMenu();
