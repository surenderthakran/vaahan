'use strict';

(function() {
  let _canvasMap;
  let _contextMap;

  let _canvasCar;
  let _contextCar;

  let _trackData;
  let _carData;

  document.addEventListener('DOMContentLoaded', function(){
    document.getElementById('track').addEventListener('change', (event) => {
      getTrack(event.target.value);
    });
    _canvasMap = document.getElementById('map');
    _canvasCar = document.getElementById('car');
  });

  function getX(point) {
    return point.x;
  }

  function getY(point) {
    return _trackData.height - point.y;
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
      _trackData = track;
      console.log(_trackData);
      initCanvas();
    })
    .catch(function(err) {
      console.error(err);
    });
  }

  function initCanvas() {
    console.log("inside initCanvas()");
    _canvasMap.height = _trackData.height;
    _canvasMap.width = _trackData.width;
    _contextMap = _canvasMap.getContext("2d");

    _canvasCar.height = _trackData.height;
    _canvasCar.width = _trackData.width;
    _contextCar = _canvasCar.getContext("2d");

    drawMap();
    initCar();
  }

  function drawMap() {
    drawBoundary();
  }

  function drawBoundary() {
    console.log("inside drawBoundary()");
    let boundary = _trackData.boundary;
    _contextMap.fillStyle = '#8ae291';
    _contextMap.fillRect(getX(boundary.top_left), getY(boundary.top_left), boundary.width, boundary.height);
  }

  function initCar() {
    var myHeaders = new Headers();
    var myRequest = new Request(
      window.location.origin + '/api/init_car?id=' + _trackData.id,
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
      _carData = car;
      console.log(_carData);
      drawCar();
    })
    .catch(function(err) {
      console.error(err);
    });
  }

  function drawCar() {
    console.log("inside drawCar()");
    _contextCar.clearRect(0, 0, _canvasCar.width, _canvasCar.height)

    _contextCar.beginPath();
    _contextCar.moveTo(getX(_carData.left_headlight), getY(_carData.left_headlight));
    _contextCar.lineTo(getX(_carData.right_headlight), getY(_carData.right_headlight));
    _contextCar.lineTo(getX(_carData.right_taillight), getY(_carData.right_taillight));
    _contextCar.lineTo(getX(_carData.left_taillight), getY(_carData.left_taillight));
    _contextCar.lineTo(getX(_carData.left_headlight), getY(_carData.left_headlight));
    _contextCar.closePath();
    _contextCar.stroke();
    _contextCar.fillStyle = "yellow";
    _contextCar.fill();

    setTimeout(updateCar, 1000);
  }

  function updateCar() {
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
      _carData = car;
      console.log(_carData);
      drawCar();
    })
    .catch(function(err) {
      console.error(err);
    });
  }
})();
