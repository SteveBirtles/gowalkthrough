function updateGamesList(id) {

    $.ajax({
        url: '/game/list/' + id,
        type: 'GET',
        success: response => {

            if (response.hasOwnProperty('error')) {

                alert(response.error);

            } else {

                $('#console').text(response.consoleName);

                let gamesHTML = `<div class="container">`
                                    + `<div class="row mb-2">`
                                    + `<div class="col-3 bg-light font-weight-bold">Game Name</div>`
                                    + `<div class="col-2 bg-light font-weight-bold">Year</div>`
                                    + `<div class="col-2 bg-light font-weight-bold">Sales</div>`
                                    + `<div class="col-3 bg-light font-weight-bold">Image</div>`
                                    + `<div class="col-2 text-right bg-light font-weight-bold">Options</div>`
                                  + `</div>`;

                for (let game of response.gamesList) {
                    gamesHTML += `<div class="row mb-2">`
                                    + `<div class="col-3">${game.name}</div>`
                                    + `<div class="col-2">${game.year}</div>`
                                    + `<div class="col-2">${game.sales}</div>`
                                    + `<div class="col-3 small"><a href="${game.imageURL}" target=”_blank”><img width="120" height="90" src="${game.imageURL}"></a></div>`
                                    + `<div class="col-2 text-right">`
                                        + `<a class="btn btn-sm btn-success"  href="/client/editgame.html?id=${game.id}">Edit</a>`
                                    +`</div>`
                                 + `</div>`;
                }
                gamesHTML += `</div>`;

                $('#games').html(gamesHTML);

            }

        }
    });

}

function pageLoad() {

    let currentPage = window.location.href;
    Cookies.set("breadcrumb", currentPage);

    let params = getQueryStringParameters();
    let id = params['id'];

    updateGamesList(id);

    $("#new").attr("href", "/client/editgame.html?id=-1&consoleId=" + id)

}