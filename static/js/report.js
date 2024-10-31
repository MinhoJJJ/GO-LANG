$(document).ready(function() {
    let chart;
    let userId=$('#userId').val();

    const currentDate = new Date();
    const currentYear = currentDate.getFullYear();
    const currentMonth = currentDate.getMonth() + 1;

    $('#year').text(currentYear);
    $('#month').text(currentMonth);

    // 초기 데이터 로드
    loadData('expense',currentMonth);

    // 지출, 수입 선택 기능
    $('.report-type-selector .btn').on('click', function() {
        $('.report-type-selector .btn').removeClass('active');
        $(this).addClass('active');
        const type = $(this).data('type');
        loadData(type,currentMonth);

    });
    // 전체내역 불러오기
    function loadData(type,month) {
        $.ajax({
            url: `/getMyhistory.do`,
            type: 'GET',
            dataType: 'json',
            data: { type: type , id: userId , month:month},
            success: function(response) {
                const processedData = processData(response);
                updateChart(processedData);
                updateList(processedData);
                updateTotalAmount(type, processedData.total);
            },
            error: function(xhr, status, error) {
                console.error("데이터 로드 중 오류 발생:", error);
            }
        });
    }

    function processData(data) {
        const total = data.reduce((sum, item) => sum + item.amount, 0);
        const colors = generateRandomColors(data.length);

        return {
            labels: data.map(item => item.category),
            values: data.map(item => item.amount),
            colors: colors,
            percentages: data.map(item => ((item.amount / total) * 100).toFixed(1)),
            total: total
        };
    }
    function updateChart(data) {
        const ctx = $('#expenseChart')[0].getContext('2d');

        if (chart) {
            chart.destroy();
        }

        chart = new Chart(ctx, {
            type: 'pie',
            data: {
                labels: data.labels,
                datasets: [{
                    data: data.values,
                    backgroundColor: data.colors
                }]
            },
            options: {
                responsive: true,
                plugins: {
                    legend: {
                        display: false
                    }
                }
            }
        });
    }

    function updateList(data) {
        const $list = $('.expense-list');
        $list.empty();

        data.labels.forEach((label, index) => {
            addExpenseItem(data.colors[index], label, data.values[index].toLocaleString(), data.percentages[index]);
        });
    }

    function addExpenseItem(color, name, amount, percentage) {
        const item = $('<li>').addClass('expense-item').html(`
            <div class="expense-info">
                <span class="expense-color" style="background-color: ${color};"></span>
                <span class="expense-name">${name}</span>
            </div>
            <span class="expense-amount">${amount}원</span>
            <span class="expense-percentage">${percentage}%</span>
        `);
        $('.expense-list').append(item);
    }

    function updateTotalAmount(type, total) {
        const totalAmount = $('.total-amount');
        console.log("durl");
        const typeText = type === 'expense' ? '지출' : '수입';
        totalAmount.text(`${typeText} ${total.toLocaleString()}원`);

        // 타입에 따라 색상 변경
        if (type === 'expense') {
            totalAmount.css('color', '#dc3545'); // 빨간색 (지출)
        } else {
            totalAmount.css('color', '#007bff'); // 파란색 (수입)
        }
    }

    // 차트 색상
    function generateRandomColors(count) {
        const colors = [];
        for (let i = 0; i < count; i++) {
            colors.push(getRandomColor());
        }
        return colors;
    }

    function getRandomColor() {
        const letters = '0123456789ABCDEF';
        let color;
        do {
            color = '#';
            for (let i = 0; i < 6; i++) {
                color += letters[Math.floor(Math.random() * 16)];
            }
        } while (isDarkColor(color));
        return color;
    }

    function isDarkColor(color) {
        const r = parseInt(color.substr(1, 2), 16);
        const g = parseInt(color.substr(3, 2), 16);
        const b = parseInt(color.substr(5, 2), 16);
        const brightness = (r * 299 + g * 587 + b * 114) / 1000;
        return brightness < 128;
    }
});