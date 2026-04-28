document.getElementById('loginButton').addEventListener('click', function() {
const emailValue = document.getElementById('emailInput').value;
const passwordValue = document.getElementById('passwordInput').value;
const messageEl = document.getElementById('message');

// プレゼン本番用のお守り（緊急回避）
// バックエンドが落ちていても、このメアドなら必ず画面を進められます！
if(emailValue === 'demo@test.com') {
location.href = 'company-list.html';
return;
}

const sendData = {
email: emailValue,
password: passwordValue
};

fetch('http://localhost:8000/api/login', { 
method: 'POST',
headers: {
'Content-Type': 'application/json'
},
body: JSON.stringify(sendData),
credentials: 'include'
})
.then(response => {
if (response.ok) {
// 成功したら一覧へ
location.href = 'company-list.html';
} else {
// 失敗したら赤文字でエラー表示
messageEl.innerText = 'メールアドレスかパスワードが違います';
messageEl.style.color = '#8f2e14'; // 弁柄色
}
})
.catch(error => {
console.error('通信エラー:', error);
messageEl.innerText = 'サーバーと通信できませんでした';
messageEl.style.color = '#8f2e14';
});
});