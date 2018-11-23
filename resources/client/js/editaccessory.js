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

    if (id !== '-1') {
        loadAccessory();
        resetDeleteButton();
    } else {
        if (params['consoleId'] !== undefined) {
            $("[name='consoleId']").val(params['consoleId']);
        }
    }

    resetForm();

}

function loadAccessory() {

    $.ajax({
        url: '/accessory/get/' + id,
        type: 'GET',
        success: accessoryDetails => {
            if (accessoryDetails.hasOwnProperty('error')) {
                alert(accessoryDetails.error);
            } else {
                $("[name='consoleId']").val(accessoryDetails.consoleId);
                $("[name='category']").val(accessoryDetails.category);
                $("[name='description']").val(accessoryDetails.description);
                $("[name='imageURL']").val(accessoryDetails.imageURL);
                $("[name='quantity']").val(accessoryDetails.quantity);
                $("[name='thirdParty']").prop("checked", accessoryDetails.thirdParty);
            }
        }
    });

}


function resetForm() {

    const form = $('#accessoryForm');

    form.unbind("submit");
    form.submit(event => {
        event.preventDefault();
        $.ajax({
            url: '/accessory/save/' + id,
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
            let r = confirm("Are you sure you want to delete this accessory?");
            if (r === true) {
                $.ajax({
                    url: '/accessory/delete/' + id,
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
