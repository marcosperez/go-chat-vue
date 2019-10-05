function ApiClient() {
    // Inicializacion de recursos de la app
    this.getChatMessage = function () {
        return new Promise((resolve, reject) => {
            var pUsers = this.getUsers();
            return Promise.all([pUsers]).then(([users]) => {
                resolve({
                    users: users
                })
            }).catch((err) => {
                console.log(err);
                reject("Todo salio mal!!!")
            });
        });
    }

    this.getUsers = function () {
        return new Promise((resolve, reject) => {
            fetch(`${config.baseUrl}/users`)
                .then((response) => response.json().then((users) => resolve(users)))
                .catch((err) => reject(err))
        });
    }

    // login de usuario....
    this.login = function (username, password) {
        return fetch(`${config.baseUrl}/users`, {
            method: 'POST',
            body: JSON.stringify({ name: username, password: password }),
            headers: {
                "Content-Type": "application/json"
            }
        }).then((response) => response.json())
    }
}

var apiClient = new ApiClient();