function sum(a, b) {
  return a + b;
}

function getRequest(host) {
  fetch(host)
    .then((response) => response.json())
    .then((data) => console.log(data));
}

module.exports = sum;
module.exports = getRequest;
