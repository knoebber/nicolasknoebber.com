const treeAPI = 'https://atmxemymx0.execute-api.us-west-2.amazonaws.com/prod/trees';

const prebuiltTrees = [
  {
    name: 'Basic',
    left: {
      length: 12,
      angle: 20,
    },
    right: {
      length: 12,
      angle: 20,
    },
  },
  {
    name: 'Grow right',
    left: {
      length: 18,
      angle: 14,
    },
    right: {
      length: 3,
      angle: 48,
    },
  },
  {
    name: 'Grow left',
    left: {
      length: 3,
      angle: 48,
    },
    right: {
      length: 18,
      angle: 14,
    },
  },
  {
    name: 'Symmetrical',
    left: {
      length: 12,
      angle: 60,
    },
    right: {
      length: 12,
      angle: 60,
    },
  },
  {
    name: 'Random',
  },
];

function rndInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1) ) + min;
}

const select = document.getElementById('prebuilt-trees');
prebuiltTrees.forEach(({ name },i) => {
  const option = document.createElement('option');
  option.value = i;
  select.appendChild(option);
  option.appendChild(document.createTextNode(name));
});

function changeTree() {
  const newTree = prebuiltTrees[parseInt(select.value, 10)];
  if (!newTree) return;
  if (newTree.name === 'Random') {
    newTree.right = { length: rndInt(-25, 25), angle: rndInt(0, 360) };
    newTree.left = { length: rndInt(-25, 25), angle: rndInt(0, 360) };
  }
  document.getElementById('right-length').value = newTree.right.length;
  document.getElementById('right-angle').value = newTree.right.angle;
  document.getElementById('left-length').value = newTree.left.length;
  document.getElementById('left-angle').value = newTree.left.angle;
  createTree(false);
}

function createTree(submitPressed) {
  const selectTree = prebuiltTrees[parseInt(select.value, 10)];
  if (submitPressed && selectTree && selectTree.name === 'Random') changeTree();
  else if (submitPressed) select.value = "-1";

  document.getElementById('create-tree-button').disabled = true;
  document.getElementById('prebuilt-trees').disabled = true;
  const rightLength = document.getElementById('right-length').value;
  const rightAngle = document.getElementById('right-angle').value;
  const leftLength = document.getElementById('left-length').value;
  const leftAngle = document.getElementById('left-angle').value;

  fetch(treeAPI, {
    method: 'POST',
    body: JSON.stringify({
      rightLength: parseInt(rightLength),
      rightAngle: parseInt(rightAngle),
      leftLength: parseInt(leftLength),
      leftAngle: parseInt(leftAngle)
    }),
    headers: {
      'Content-Type':'application/json'
    }
  }).then(response => response.json())
    .then(data => document.getElementById('lambda-go-tree').src = `https://nicolasknoebber.com/posts/images/trees/${data.message}`)
    .catch(err => console.log(err))
    .finally(() => {
      document.getElementById('create-tree-button').disabled = false;
      document.getElementById('prebuilt-trees').disabled = false;
    });
}

