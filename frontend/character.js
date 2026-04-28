// つぶやくセリフのリスト
const messages = [
"今日もえらい！", "眉毛キマってるね！", "一歩進んで二歩下がる～",
"第二新卒で頑張るか！", "国外逃亡も視野に...", "就活は婚活！最後はフィーリングや！",
"貴社のご発展をお祈り申し上げます（皮肉）", "市場価値アゲアゲ～！",
"美味しい企業ないかな〜", "適度に休むのも大事だよ！",
"面接官もお疲れさまやなぁ", "落ちたのは縁がなかっただけ！",
"モンスターでも飲んで元気出せや"
];

// 食べた時の専用セリフリスト
const eatMessages = [
"モグモグ…この企業、歯ごたえある！",
"お祈りメールは栄養満点や！",
"こんな企業、こっちから願い下げやで！ムシャァ！",
"私の良さがわからんとはお気の毒に！モグモグ！",
"はい次！美味しいごはんでした！",
"落ちた数だけ強くなる！ゴクン！"
];

// 【追加】内定した時の専用お祝いセリフリスト
const celebrateMessages = [
"内定おめでとうー！！！本当に、本当によく頑張ったね！",
"すごすぎる！これまでの努力が実を結んだね！",
"やったー！！最高の結果だね！！",
"あなたの魅力がしっかり伝わったんだね！おめでとう！"
];

// 画面が読み込まれた瞬間に動く設定
window.onload = function() {
startRandomSpeech();
updateLevelDisplay();
checkStatusChange();
};

function startRandomSpeech() {
setInterval(() => {
const bubbleText = document.getElementById('globalMessageText');
if (!bubbleText) return; 
const randomIndex = Math.floor(Math.random() * messages.length);
bubbleText.innerText = messages[randomIndex];
}, 7000);
}

function updateLevelDisplay() {
const levelText = document.getElementById('levelNumber');
if (!levelText) return;
const savedLevel = localStorage.getItem('rejectionLevel') || 0;
levelText.innerText = savedLevel;
}

// 戻ってきた時に「不合格」か「内定」かを判定する
function checkStatusChange() {
    if (localStorage.getItem('triggerEat') === 'true') {
        setTimeout(() => {
            eatCompanyAction();
            localStorage.removeItem('triggerEat');
        }, 500); 
    } else if (localStorage.getItem('triggerCelebrate') === 'true') {
        setTimeout(() => {
            celebrateAction();
            localStorage.removeItem('triggerCelebrate');
        }, 500);
    }
}

function eatCompanyAction() {
const charImg = document.getElementById('globalCharImg');
const bubbleText = document.getElementById('globalMessageText');

if (!charImg || !bubbleText) return;

const randomEatIndex = Math.floor(Math.random() * eatMessages.length);
bubbleText.innerText = eatMessages[randomEatIndex];

charImg.classList.remove('eat-grow-anim');
void charImg.offsetWidth; 
charImg.classList.add('eat-grow-anim');

let currentLevel = parseInt(localStorage.getItem('rejectionLevel') || 0);
currentLevel++;
localStorage.setItem('rejectionLevel', currentLevel);

updateLevelDisplay();
}

// 【追加】内定のお祝いをする機能
function celebrateAction() {
    const bubbleText = document.getElementById('globalMessageText');
    if (!bubbleText) return;

    // お祝いセリフをランダムに選んで表示
    const randomCelIndex = Math.floor(Math.random() * celebrateMessages.length);
    bubbleText.innerText = celebrateMessages[randomCelIndex];
}