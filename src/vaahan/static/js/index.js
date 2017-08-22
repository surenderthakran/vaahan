'use strict';

(function() {
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

  document.addEventListener('DOMContentLoaded', function(){
    document.getElementById('track').addEventListener('change', (event) => {
      getTrack(event.target.value);
    });
    _map.canvas = document.getElementById('map');
    _car.canvas = document.getElementById('car');
  });

  function getX(point) {
    return point.x;
  }

  function getY(point) {
    return _map.track.height - point.y;
  }

  function getTrack(trackId) {
    var myHeaders = new Headers();
    var myRequest = new Request(
      window.location.origin + '/api/get_track?id=' + trackId,
      {
        method: 'GET',
        headers: myHeaders,
        mode: 'cors',
        cache: 'default',
      }
    );

    fetch(myRequest)
    .then(function(response) {
      console.log(response);
      if(response.ok) {
        return response.json();
      } else {
        throw Error(response.statusText);
      }
    })
    .then(function(track) {
      _map.track = track;
      console.log(_map.track);
      initCanvas();
    })
    .catch(function(err) {
      console.error(err);
    });
  }

  function initCanvas() {
    console.log("inside initCanvas()");
    _map.canvas.height = _map.track.height;
    _map.canvas.width = _map.track.width;
    _map.context = _map.canvas.getContext("2d");

    _car.canvas.height = _map.track.height;
    _car.canvas.width = _map.track.width;
    _car.context = _car.canvas.getContext("2d");

    drawMap();
    initCar();
  }

  function drawMap() {
    drawBoundary();
  }

  function drawBoundary() {
    console.log("inside drawBoundary()");
    let boundary = _map.track.boundary;
    _map.context.fillStyle = '#8ae291';
    _map.context.fillRect(getX(boundary.top_left), getY(boundary.top_left), boundary.width, boundary.height);
  }

  function initCar() {
    var myHeaders = new Headers();
    var myRequest = new Request(
      window.location.origin + '/api/init_car?id=' + _map.track.id,
      {
        method: 'GET',
        headers: myHeaders,
        mode: 'cors',
        cache: 'default',
      }
    );

    _car.runUpdateLoop = false;
    document.getElementById('start-pause').textContent = "Start Driving";

    fetch(myRequest)
    .then(function(response) {
      console.log(response);
      if(response.ok) {
        return response.json();
      } else {
        throw Error(response.statusText);
      }
    })
    .then(function(car) {
      _car.data = car;
      console.log(_car.data);
      initCarControls();
      drawCar();
    })
    .catch(function(err) {
      console.error(err);
    });
  }

  function initCarControls() {
    document.getElementById('start-pause').addEventListener('click', (event) => {
      if (_car.data.status === "stopped") {
        driveCar();
        event.target.textContent = "Pause Driving";
      } else if (_car.data.status === "driving") {
        pauseCar();
        event.target.textContent = "Resume Driving";
      } else if (_car.data.status === "paused") {
        driveCar();
        event.target.textContent = "Pause Driving";
      }
    });

    document.getElementById('reset-car').addEventListener('click', (event) => {
      initCar();
      document.getElementById('start-pause').value = "Start Driving";
    });
  }

  function drawCar() {
    console.log("inside drawCar()");
    _car.context.clearRect(0, 0, _car.canvas.width, _car.canvas.height)

    _car.context.beginPath();
    _car.context.moveTo(getX(_car.data.left_headlight), getY(_car.data.left_headlight));
    _car.context.lineTo(getX(_car.data.right_headlight), getY(_car.data.right_headlight));
    _car.context.lineWidth = 3;
    _car.context.stroke();
    _car.context.lineTo(getX(_car.data.right_taillight), getY(_car.data.right_taillight));
    _car.context.lineWidth = 0.5;
    _car.context.lineTo(getX(_car.data.left_taillight), getY(_car.data.left_taillight));
    _car.context.lineTo(getX(_car.data.left_headlight), getY(_car.data.left_headlight));
    _car.context.closePath();
    _car.context.stroke();
    _car.context.fillStyle = "yellow";
    _car.context.fill();

    _car.context.moveTo(getX(_car.data.back_center), getY(_car.data.back_center));
    _car.context.lineTo(getX(_car.data.front_center), getY(_car.data.front_center));
    _car.context.stroke();

    _car.context.fillStyle = "red";
    _car.context.fillRect(getX(_car.data.left_headlight) - 2, getY(_car.data.left_headlight) - 2, 4, 4);

    _car.context.fillStyle = "blue";
    _car.context.fillRect(getX(_car.data.right_headlight) - 2, getY(_car.data.right_headlight) - 2, 4, 4);

    _car.context.fillStyle = "red";
    _car.context.fillRect(getX(_car.data.left_taillight) - 2, getY(_car.data.left_taillight) - 2, 4, 4);

    _car.context.fillStyle = "blue";
    _car.context.fillRect(getX(_car.data.right_taillight) - 2, getY(_car.data.right_taillight) - 2, 4, 4);

    if (_car.runUpdateLoop) {
      setTimeout(updateCarData, 1000);
    }
  }

  function driveCar() {
    var myHeaders = new Headers();
    var myRequest = new Request(
      window.location.origin + '/api/drive_car',
      {
        method: 'GET',
        headers: myHeaders,
        mode: 'cors',
        cache: 'default',
      }
    );

    _car.runUpdateLoop = true;

    fetch(myRequest)
    .then(function(response) {
      console.log(response);
      if(response.ok) {
        return response.json();
      } else {
        throw Error(response.statusText);
      }
    })
    .then(function(car) {
      _car.data = car;
      console.log(_car.data);
      drawCar();
      document.getElementById('start-pause').textContent = "Pause Driving";
    })
    .catch(function(err) {
      console.error(err);
    });
  }

  function updateCarData() {
    var myHeaders = new Headers();
    var myRequest = new Request(
      window.location.origin + '/api/get_car',
      {
        method: 'GET',
        headers: myHeaders,
        mode: 'cors',
        cache: 'default',
      }
    );

    fetch(myRequest)
    .then(function(response) {
      console.log(response);
      if(response.ok) {
        return response.json();
      } else {
        throw Error(response.statusText);
      }
    })
    .then(function(car) {
      _car.data = car;
      console.log(_car.data);
      drawCar();
    })
    .catch(function(err) {
      console.error(err);
    });
  }
})();
