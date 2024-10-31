$(document).ready(function() {

    // 챗관련
    const chatIcon = $('#chat-icon');
    const chatContainer = $('#chat-container');
    const chatMessages = $('#chat-messages');
    const userInput = $('#user-input');
    const sendButton = $('#send-button');

    chatIcon.click(function() {
        chatContainer.toggle();
    });

    function addMessage(message, isUser = false) {
        const messageElement = $('<div>').addClass('message');
        if (isUser) {
            messageElement.addClass('user-message');
        } else {
            messageElement.addClass('ai-message');
        }
        messageElement.text(message);
        chatMessages.append(messageElement);
        chatMessages.scrollTop(chatMessages[0].scrollHeight);
    }

    function sendMessage() {
        const message = userInput.val().trim();
        if (message) {
            addMessage(message, true);
            userInput.val('');

            // API 호출
            $.ajax({
                url: 'http://localhost:11434/api/generate',
                type: 'POST',
                contentType: 'application/json',
                data: JSON.stringify({
                    model: 'llama3.1:8b',
                    prompt: message
                }),
                success: function(response) {
                    let fullResponse = '';
                    response.split('\n').forEach(line => {
                        if (line) {
                            const jsonResponse = JSON.parse(line);
                            if (jsonResponse.response) {
                                fullResponse += jsonResponse.response;
                            }
                        }
                    });
                    addMessage(fullResponse);
                },
                error: function(jqXHR, textStatus, errorThrown) {
                    console.error("Error:", textStatus, errorThrown);
                    addMessage("죄송합니다. 오류가 발생했습니다.");
                }
            });
        }
    }

    sendButton.click(sendMessage);
    userInput.keypress(function(e) {
        if (e.which == 13) {
            sendMessage();
            return false;
        }
    });

    //페이징 관련
    var currentPage = window.location.pathname.split('/').pop().split('.')[0];

    if (document.baseURI.includes(currentPage)) {
        $('#'+currentPage).addClass('active');
        $('#'+currentPage).css('font-weight', '900');
    } else {
        $(this).removeClass('active');
    }

    // 초창기 코드
    function generateText() {

        $.ajax({
            url: 'http://localhost:11434/api/generate',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                model: llamaModel,
                prompt: prompt
            }),
            success: function(response) {
                let fullResponse = '';
                response.split('\n').forEach(line => {
                    if (line) {
                        const jsonResponse = JSON.parse(line);
                        if (jsonResponse.response) {
                            fullResponse += jsonResponse.response;
                        }
                    }
                });
                console.log(fullResponse);
            },
            error: function(xhr, textStatus, error) {
                if (xhr.status === 0) {
                    console.error("오류로 인해 로컬 개발서버를 찾지 못하였습니다. 잠시 후 다시 시도해 주세요.");
                } else {
                    console.error(xhr.responseText);
                    console.error(error);
                }
            }
        });
    }

    // generateText();
});