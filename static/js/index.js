jQuery(function($) {
    $("#retrieve_balance_sheet_btn").click(function () {
        let businessName = $("#business_name").val();
        let provider = $("#accounting_provider").val();
        let data = {
            business_name: businessName,
            accounting_provider: provider,
        };
        console.log(data);

        $.ajax({
            url: "api/retrieve_balance_sheet",
            method: "POST",
            data: JSON.stringify(data),
            dataType: "json",
            contentType: "application/json; charset=utf-8",
        }).done(function(resp) {
            console.log(resp);
            let entries = resp["data"];
            let table = "<table style='table-layout: fixed; width:100%;'><thead><th>Year</th><th>Month</th><th>Profit/Loss</th><th>Assets Value</th></<thead>";
            for (let i = 0; i < entries.length; i++) {
                table += `<tbody><td>${entries[i]["year"]}</td><td>${entries[i]["month"]}</td><td>$${entries[i]["profit_or_loss"]}</td><td>$${entries[i]["assets_value"]}</td></tbody>`;
            }
            table += "</table>";
            $("#balance_sheet_display").html(table);
        }).fail(function(resp) {
            alert(resp["responseJSON"]["error_message"]);
            console.log(resp["error_message"]);
        });
    });

    $("#submit_loan_application_btn").click(function () {
        let businessName = $("#business_name").val();
        let yearEstablished = $("#year_established").val();
        let loanAmount = $("#loan_amount").val();
        let provider = $("#accounting_provider").val();
        let data = {
            business_name: businessName,
            year_established: parseInt(yearEstablished),
            loan_amount: parseInt(loanAmount),
            accounting_provider: provider,
        };
        console.log(data);

        $.ajax({
            url: "api/submit_loan_application",
            method: "POST",
            data: JSON.stringify(data),
            dataType: "json",
            contentType: "application/json; charset=utf-8",
        }).done(function(resp) {
            console.log(resp);
            let verdict = resp['data']['verdict'];
            if (verdict === true) {
                let preAssessmentValue = resp['data']['pre_assessment_value'];
                let eligibleLoanAmount = resp['data']['eligible_loan_amount'];
                $("#loan_application_result").html(`Your loan application is successful. PreAssessment value = ${preAssessmentValue}, Eligible loan amount = $${eligibleLoanAmount}`);
            } else {
                $("#loan_application_result").html("Your loan application is unsuccessful.");
            }
        }).fail(function(resp) {
            alert(resp["responseJSON"]["error_message"]);
            console.log(resp);
        });
    });
});
