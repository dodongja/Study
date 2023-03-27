const quotes = [
    {
        quote: "외롭노",
        author: "최강혁",
    },
    {
        quote: "나도 외롭노",
        author: "이동혁",
    },
    {
        quote: "오늘부터 자기개발뿐이다!",
        author: "이동혁, 최강혁",
    },
    {
        quote: "행복하네요 ㅎㅎ",
        author: "이유찬",
    },
    {
        quote: "최강혁 많이 죽었네",
        author: "이유찬",
    },
    {
        quote: "많이컸노 유차이",
        author: "최강혁",
    },
]

const quote = document.querySelector("#quote span:first-child");
const author = document.querySelector("#quote span:last-child");

const todayQoute = quotes[Math.floor(Math.random() * quotes.length)];

quote.innerText = todayQoute.quote;
author.innerText = todayQoute.author;