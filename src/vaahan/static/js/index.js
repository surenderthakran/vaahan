'use strict';

let _tracks = [];

let _map = {
  canvas: null,
  context: null,
  track: null,
};

let _car = {
  canvas: null,
  context: null,
  data: null,
  runUpdateLoop: false,
};

let restartOnCollision = false;

document.addEventListener('DOMContentLoaded', function(){
  _map.canvas = document.getElementById('map');
  _car.canvas = document.getElementById('car');
  getCurrentSimulation();
  initCarControls();
});

function initCarControls() {
  document.getElementById('start-pause').addEventListener('click', (event) => {
    console.log(_car.data.status);
    if (_car.data.status === "STOP") {
      driveCar();
    } else if (_car.data.status === "DRIVE") {
      stopCar();
    }
  });

  document.getElementById('reset-car').addEventListener('click', (event) => {
    initCar();
    document.getElementById('start-pause').value = "Start Driving";
  });

  document.getElementById('restart').addEventListener('change', (event) => {
    restartOnCollision = event.target.checked;
    updateRestartConf(restartOnCollision);
  });
}

function populateTracksDropDown() {
  var tracksSelect = document.getElementById("tracks");
  for (var i = 0, len = _tracks.length; i < len; i++) {
    var option = document.createElement('option');
    option.value = _tracks[i].id;
    option.text = _tracks[i].name;
    tracksSelect.add(option);
  }
}

function initCanvas() {
  console.log("inside initCanvas()");
  _map.canvas.height = _map.track.height;
  _map.canvas.width = _map.track.width;
  _map.context = _map.canvas.getContext("2d");

  _car.canvas.height = _map.track.height;
  _car.canvas.width = _map.track.width;
  _car.context = _car.canvas.getContext("2d");
}

function drawMap() {
  document.getElementById('track').textContent = _map.track.id + ". " + _map.track.name;
  drawBoundary();
}

function updateCar() {
  restartOnCollision = _car.data.restartOnCollision;
  document.getElementById('restart').checked = restartOnCollision;

  if (_car.data.status === "STOP") {
    _car.runUpdateLoop = false;
    document.getElementById('start-pause').textContent = "Start Driving";
    document.getElementById('status').textContent = "Car has stopped.";
  } else if (_car.data.status === "DRIVE") {
    _car.runUpdateLoop = true;
    document.getElementById('start-pause').textContent = "Stop Driving";
    document.getElementById('status').textContent = "Car is driving.";
  } else if (_car.data.status === "COLLISION") {
    _car.runUpdateLoop = restartOnCollision;
    if (restartOnCollision) {
      document.getElementById('start-pause').textContent = "Stop Driving";
    } else {
      document.getElementById('start-pause').textContent = "Start Driving";
    }
    document.getElementById('status').textContent = "Car has collided.";
  }

  drawCar();
  if (_car.runUpdateLoop) {
    setTimeout(getCar, 250);
  }
}
