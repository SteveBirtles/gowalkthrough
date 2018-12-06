let id = -1;

function pageLoad() {

    let lastPage =  Cookies.get("breadcrumb");
    $("#back").attr("href", lastPage);

    let currentPage = window.location.href;
    Cookies.set("destination", currentPage);

    checkLogin();

    $("#logout").click(event => {
        Cookies.remove("sessionToken");
        window.location.href = "/client/index.html";
    });

    let params = getQueryStringParameters();
    if (params['id'] !== undefined) {
        id = parseInt(params['id']);
    }

    if (id !== -1) {
        loadGame();
        resetDeleteButton();
    } else {
        if (params['consoleId'] !== undefined) {
            $("[name='consoleId']").val(params['consoleId']);
        }
    }

    resetForm();

}

function loadGame() {

    $.ajax({
        url: '/game/get/' + id,
        type: 'GET',
        success: gameDetails => {
            if (gameDetails.hasOwnProperty('error')) {
                alert(gameDetails.error);
            } else {
                $("[name='consoleId']").val(gameDetails.consoleId);
                $("[name='name']").val(gameDetails.name);
                $("[name='year']").val(gameDetails.year);
                $("[name='sales']").val(gameDetails.sales);
                $("[name='imageURL']").val(gameDetails.imageURL);
            }
        }
    });

}

function resetForm() {

    const form = $('#gameForm');

    form.unbind("submit");
    form.submit(event => {
        event.preventDefault();
        $.ajax({
            url: '/game/save/' + id,
            type: 'POST',
            data: form.serialize(),
            success: response => {
                if (response === 'OK') {
                    window.location.href = $("#back").attr("href");
                } else {
                    alert(response);
                }
            }
        });
    });
}

function resetDeleteButton() {

    $('#delete')
        .css('visibility', 'visible')
        .click(event => {
            let r = confirm("Are you sure you want to delete this game?");
            if (r === true) {
                $.ajax({
                    url: '/game/delete/' + id,
                    type: 'POST',
                    success: response => {
                        if (response === 'OK') {
                            window.location.href = $("#back").attr("href");
                        } else {
                            alert(response);
                        }
                    }
                });
            }
        }
    );

}
