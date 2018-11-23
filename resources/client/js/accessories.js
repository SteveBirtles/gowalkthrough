function updateAccessoriesList(id) {

    $.ajax({
        url: '/accessory/list/' + id,
        type: 'GET',
        success: response => {

            if (response.hasOwnProperty('error')) {

                alert(response.error);

            } else {

                $('#console').text(response.consoleName);

                let accessoriesHTML = `<div class="container">`
                    + `<div class="row mb-2">`
                    + `<div class="col-3 bg-light font-weight-bold">Description</div>`
                    + `<div class="col-2 bg-light font-weight-bold">Quantity</div>`
                    + `<div class="col-2 bg-light font-weight-bold">ThirdParty</div>`
                    + `<div class="col-3 bg-light font-weight-bold">Image</div>`
                    + `<div class="col-2 text-right bg-light font-weight-bold">Options</div>`
                    + `</div>`;

                for (let accessory of response.accessoriesList) {
                    accessoriesHTML += `<div class="row mb-2">`
                        + `<div class="col-3">${accessory.description}</div>`
                        + `<div class="col-2">${accessory.quantity}</div>`
                        + `<div class="col-2">${accessory.thirdParty}</div>`
                        + `<div class="col-3 small"><a href="${accessory.imageURL}" target=”_blank”><img width="120" height="90" src="${accessory.imageURL}"></a></div>`
                        + `<div class="col-2 text-right">`
                        + `<a class="btn btn-sm btn-success"  href="/client/editaccessory.html?id=${accessory.id}">Edit</a>`
                        +`</div>`
                        + `</div>`;
                }
                accessoriesHTML += `</div>`;

                $('#accessories').html(accessoriesHTML);

            }

        }
    });

}


function pageLoad() {

    let currentPage = window.location.href;
    Cookies.set("breadcrumb", currentPage);

    let params = getQueryStringParameters();
    let id = params['id'];

    updateAccessoriesList(id);

    $("#new").attr("href", "/client/editaccessory.html?id=-1&consoleId=" + id)

}