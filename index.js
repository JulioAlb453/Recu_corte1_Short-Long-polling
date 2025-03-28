let requestCount = 0;
const maxRequests = 10;
const users = [];
let countHombres = 0;
let countMujeres = 0;
const requestTimeout = 5000;

document.getElementById('user-form').addEventListener('submit', function (event) {
    event.preventDefault();
    const nombre = document.getElementById('nombre').value;
    const edad = document.getElementById('edad').value;
    const genero = document.getElementById('genero').value;

    const newUser = { nombre, edad, genero };

    fetch('http://localhost:8080/addPerson', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(newUser),
    })
        .then(response => response.json())
        .then(data => {
            console.log('Usuario agregado:', data);
            users.push(newUser);
            if (genero === 'Hombre') {
                countHombres++;
            } else {
                countMujeres++;
            }
            updateUserList();
            updateGenderCount();
        })
        .catch(error => {
            console.error('Error al agregar usuario:', error);
        });

    document.getElementById('nombre').value = '';
    document.getElementById('edad').value = '';
    document.getElementById('genero').value = 'Hombre';
});

function getUsersShortPolling() {
    if (requestCount >= maxRequests) {
        document.getElementById('error-message').textContent = 'Límite de peticiones alcanzado. Intenta de nuevo más tarde.';
        document.getElementById('error-message').style.display = 'block';
        return;
    }

    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), requestTimeout);

    fetch('http://localhost:8080/NewPersonAdded', { signal: controller.signal })
        .then(response => response.json())
        .then(data => {
            clearTimeout(timeoutId);
            if (!data || data.length === 0) {
                return;
            }
            updateUserList(data);
            requestCount++;
        })
        .catch(error => {
            if (error.name === 'AbortError') {
                console.error('La petición fue abortada por timeout');
            } else {
                console.error('Error en Short Polling:', error);
            }
        });
}

function getGenderCountLongPolling() {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), requestTimeout);

    fetch('http://localhost:8080/CountGender', { signal: controller.signal })
        .then(response => response.json())
        .then(data => {
            clearTimeout(timeoutId);
            if (!data || !data.hombres || !data.mujeres) {
                return;
            }
            countHombres = data.hombres;
            countMujeres = data.mujeres;
            updateGenderCount();
        })
        .catch(error => {
            if (error.name === 'AbortError') {
                console.error('La petición fue abortada por timeout');
            } else {
                console.error('Error en Long Polling:', error);
            }
        });
}
