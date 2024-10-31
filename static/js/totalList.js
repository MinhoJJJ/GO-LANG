$(document).ready(function() {
    let grid;
    let transactions; // 전역 변수로 선언
    let userId=$('#userId').val();
    // AJAX 요청으로 데이터 가져오기
    $.ajax({
        url: '/totalListData.do',
        type: 'GET',
        data: { id: userId },
        success: function(response) {
            // 서버에서 받은 데이터 처리
            transactions = response.map((item, index) => ({
                list_ndx: index + 1,
                date: item.date,
                description: item.description,
                category: item.category,
                amount: formatAmount(item.amount),
                type: item.type === 'E' ? '지출' : '수입'
            }));

            // 그리드 초기화 및 데이터 설정
            initGrid(); // AJAX 요청이 성공한 후에 initGrid 호출
        },
        error: function(xhr, status, error) {
            console.error("데이터를 가져오는 중 오류 발생:", error);
        }
    });

    function formatAmount(amount) {
        return '₩' + amount.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    }

    function initGrid() {
        const el = document.getElementById('grid');

        grid = new tui.Grid({
            el: el,
            data: transactions,
            // ... (나머지 코드는 그대로 유지)
        });
    }
    function formatAmount(amount) {
        return '₩' + amount.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
    }

    function initGrid() {
        const el = document.getElementById('grid');

        grid = new tui.Grid({
            el: el,
            data: transactions,
            presetName: 'clean',  // Clean Theme 적용
            columns: [
                {
                    header: '날짜',
                    name: 'date',
                    width: 120,
                    filter: 'date'
                },
                {
                    header: '설명',
                    name: 'description',
                    width: 688,
                    filter: 'text'
                },
                {
                    header: '카테고리',
                    name: 'category',
                    width: 170,
                    filter: 'text'
                },
                {
                    header: '금액',
                    name: 'amount',
                    width: 140,
                    filter: 'number'
                },
                {
                    header: '거래 구분',
                    name: 'type',
                    width: 120,
                    filter: 'text',
                    align: 'center',
                    formatter: function(cell) {
                        if (cell.value === '지출') {
                            return '<span style="color: red;">' + cell.value + '</span>';
                        } else if (cell.value === '수입') {
                            return '<span style="color: blue;">' + cell.value + '</span>';
                        }
                        return cell.value;
                    }
                }
            ],
            bodyHeight: 'auto',
            columnOptions: {
                resizable: true
            },
            width: 'auto',
            rowHeaders: ['checkbox'],
            treeColumnOptions: {
                name: 'description',
                useCascadingCheckbox: true
            }
        });
    }

    function searchTransactions() {
        const startDate = $('#start-date').val();
        const endDate = $('#end-date').val();
        console.log('Searching transactions:', startDate, endDate);
        // 여기에 검색 로직을 구현하세요
    }

    function deleteSelectedTransactions() {
        const checkedRows = grid.getCheckedRows();
        if (checkedRows.length === 0) {
            alert('삭제할 항목을 선택해주세요.');
            return;
        }

        const idsToDelete = checkedRows.map(row => row.id);
        console.log('Deleting transactions:', idsToDelete);

        // 여기에 실제 삭제 로직을 구현하세요
        // 예: 서버에 삭제 요청을 보내고, 성공하면 그리드에서 행을 제거합니다.
        idsToDelete.forEach(id => {
            const rowKey = grid.getRowById(id).rowKey;
            grid.removeRow(rowKey);
        });
    }

    // 이벤트 리스너 설정
    $('#searchForm').on('submit', function(e) {
        e.preventDefault();
        searchTransactions();
    });

    // 삭제 버튼 추가 및 이벤트 리스너 설정
    $('<button id="deleteButton" class="btn btn-danger mt-3">선택 항목 삭제</button>').insertBefore('#grid');
    $('#deleteButton').on('click', deleteSelectedTransactions);
});