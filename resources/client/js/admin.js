function resetLoginForm() {

    if (Cookies.get("destination") === undefined) {
        window.location.href = "/client/index.html";
    }

    const loginForm = $('#loginForm');
    loginForm.submit(event => {
        event.preventDefault();
        $.ajax({
            url: '/admin/login',
            type: 'POST',
            data: loginForm.serialize(),
            success: response => {
                if (response.startsWith('Error:')) {
                    alert(response);
                } else {
                    Cookies.set("sessionToken", response);
                    window.location.href = Cookies.get("destination");
                }
            }
        });
    });

}

function checkLogin() {

    let currentPage = window.location.pathname;
    let token = Cookies.get("sessionToken");

    if (token !== undefined) {
        $.ajax({
            url: '/admin/check',
            type: 'GET',
            success: username => {
                if (username === "") {
                    if (currentPage !== '/client/login.html') {
                        window.location.href = '/client/login.html';
                    }
                }
            }
        });
    } else {
        if (currentPage !== '/client/login.html') {
            window.location.href = '/client/login.html';
        }
    }
}

