let currencies = [];

window.addEventListener("load",  () => {
    loadCurrencies();
    const form = document.getElementById("currencyForm")
    const fromCurrency = document.querySelector('select[name="fromCurrency"]')
    const toCurrency = document.querySelector('select[name="toCurrency"]')
    form.addEventListener("submit", (event) => {
        event.preventDefault()
        getRate(form);
    })

    fromCurrency.addEventListener("change", (event) => {
        event.preventDefault()
        getRate(form);
    })

    toCurrency.addEventListener("change", (event) => {
        event.preventDefault()
        getRate(form);
    })

})

function loadCurrencies(){
    fetch("./public/js/currency.json")
        .then(res => res.json())
        .then(data => {
            if (data) {
                currencies = data
            }
        });
}

function fetchCountryInfo(code) {
    return currencies.find(currency =>  currency.code === code ) || {}
}
function formatter(amount, currency) {
     let country = fetchCountryInfo(currency)
     let value = parseFloat(amount).toLocaleString(country.locale, {
        style: 'currency',
        currency,
    })

    return `${value} ${country.name}`
}

function getRate(form)
{
    const formData = new FormData( form );
    const request = {
        method: "POST",
        body: formData,
    };
    fetch("./api/v1/get-rate", request)
        .then((response) => response.json())
        .then((data) => {
            if (data) {
                let amount = formData.get('amount');
                let from = formData.get('fromCurrency')
                let to = formData.get('toCurrency')

                document.getElementById("fromValue").innerHTML =
                    `${formatter(amount, from)} =`
                document.getElementById("rate").innerHTML = `${formatter(data, to)}`
                console.log('response',data)
            } else {
            }
        });
}