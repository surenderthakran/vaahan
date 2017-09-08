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
    initCarControls();
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
    _map.context.fillRect(getX(boundary.shape.top_left), getY(boundary.shape.top_left), boundary.shape.width, boundary.shape.height);
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
      _car.context.clearRect(0, 0, _car.canvas.width, _car.canvas.height)
      updateCarControls();
    })
    .catch(function(err) {
      console.error(err);
    });
  }

  function initCarControls() {
    document.getElementById('start-pause').addEventListener('click', (event) => {
      console.log(_car.data.status);
      if (_car.data.status === "STOP") {
        driveCar();
        // event.target.textContent = "Stop Driving";
      } else if (_car.data.status === "DRIVE") {
        stopCar();
        // event.target.textContent = "Start Driving";
      }
    });

    document.getElementById('reset-car').addEventListener('click', (event) => {
      initCar();
      document.getElementById('start-pause').value = "Start Driving";
    });
  }

  function updateCarControls() {
    if (_car.data.status === "STOP") {
      _car.runUpdateLoop = false;
      drawCar();
      document.getElementById('start-pause').textContent = "Start Driving";
    } else if (_car.data.status === "DRIVE") {
      _car.runUpdateLoop = true;
      drawCar();
      document.getElementById('start-pause').textContent = "Stop Driving";
    }
  }

  function drawCar() {
    console.log("inside drawCar()");
    _car.context.clearRect(0, 0, _car.canvas.width, _car.canvas.height)

    _car.context.lineWidth = 0.5;

    // draw car's border.
    _car.context.beginPath();
    _car.context.moveTo(getX(_car.data.points.front_left), getY(_car.data.points.front_left));
    _car.context.lineTo(getX(_car.data.points.front_right), getY(_car.data.points.front_right));
    _car.context.lineTo(getX(_car.data.points.back_right), getY(_car.data.points.back_right));
    _car.context.lineTo(getX(_car.data.points.back_left), getY(_car.data.points.back_left));
    _car.context.lineTo(getX(_car.data.points.front_left), getY(_car.data.points.front_left));
    _car.context.closePath();
    _car.context.stroke();
    _car.context.fillStyle = "yellow";
    _car.context.fill();

    // draw color points at car's corners.
    _car.context.fillStyle = "red";
    _car.context.fillRect(getX(_car.data.points.front_left) - 2, getY(_car.data.points.front_left) - 2, 4, 4);
    _car.context.fillRect(getX(_car.data.points.back_left) - 2, getY(_car.data.points.back_left) - 2, 4, 4);

    _car.context.fillStyle = "blue";
    _car.context.fillRect(getX(_car.data.points.front_right) - 2, getY(_car.data.points.front_right) - 2, 4, 4);
    _car.context.fillRect(getX(_car.data.points.back_right) - 2, getY(_car.data.points.back_right) - 2, 4, 4);

    // draw car's sensors vision.
    _car.context.lineWidth = 0.1;
    for (var i in _car.data.sensors) {
      if (_car.data.sensors[i].intersection !== null) {
        _car.context.moveTo(getX(_car.data.sensors[i].ray.start_point), getY(_car.data.sensors[i].ray.start_point));
        _car.context.lineTo(getX(_car.data.sensors[i].intersection), getY(_car.data.sensors[i].intersection));
        _car.context.stroke();
      }
    }

    if (_car.runUpdateLoop) {
      setTimeout(getCarData, 250);
    }
  }

  function getCarData() {
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
      updateCarControls();
    })
    .catch(function(err) {
      console.error(err);
    });
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
      updateCarControls();
    })
    .catch(function(err) {
      console.error(err);
    });
  }

  function stopCar() {
    var myHeaders = new Headers();
    var myRequest = new Request(
      window.location.origin + '/api/stop_car',
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
      updateCarControls();
    })
    .catch(function(err) {
      console.error(err);
    });
  }
})();
