// web/main.js

document.addEventListener('DOMContentLoaded', function() {
    var priceTable = document.getElementById('priceTable');

    // Create a WebSocket connection to the server
    var socket = new WebSocket('ws://localhost:8080/ws');

    // Handle messages received from the server
    socket.onmessage = function(event) {
        var processedData = JSON.parse(event.data);

        console.log(processedData);

        // Update the price table
        updatePriceTable(processedData);
    };

    // Function to update the price table with the received processed data
    function updatePriceTable(processedData) {
        // Clear the existing table rows

        while (priceTable.rows.length > 1) {
            priceTable.deleteRow(1);
        }

        // Add new rows with the updated processed data
        for (var i = 0; i < processedData.length; i++) {
            var currency = processedData[i].Currency;
            var price = processedData[i].Price.toFixed(2);
            var timestamp = processedData[i].Timestamp;

            var row = priceTable.insertRow(-1);
            var currencyCell = row.insertCell(0);
            var priceCell = row.insertCell(1);
            var timestampCell = row.insertCell(2);


            currencyCell.innerHTML = currency;
            priceCell.innerHTML = price;
            timestampCell.innerHTML = timestamp;
        }
    }
});
