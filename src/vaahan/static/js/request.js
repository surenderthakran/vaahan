'use strict'

function getCurrentSimulation() {
  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/get_sim',
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
  .then(function(sim) {
    console.log(sim);
    _map.track = sim.track;
    _car.data = sim.car;
    initCanvas();
    drawMap();
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}

function getCar() {
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
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}

function initCar() {
  if (_map.track == undefined) {
    document.getElementById('status').textContent = "Select a track first!";
    return;
  }

  var myHeaders = new Headers();
  var myRequest = new Request(
    window.location.origin + '/api/init_car',
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
    updateCar();
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
    updateCar();
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
    updateCar();
  })
  .catch(function(err) {
    console.error(err);
  });
}
