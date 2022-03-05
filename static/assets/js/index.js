let rText = document.getElementById("r-text")
let hText = document.getElementById("h-text")
let kText = document.getElementById("k-text")

let clipInfo = document.getElementById("clipboard-info")

const clipPopupTime = 1500


rText.addEventListener("input", () => {
    hText.value = wanakana.toHiragana(rText.value)
    kText.value = wanakana.toKatakana(rText.value)
})

hText.addEventListener("input", () => {
    rText.value = wanakana.toRomaji(hText.value)
    kText.value = wanakana.toKatakana(hText.value)
})

kText.addEventListener("input", () => {
    hText.value = wanakana.toHiragana(kText.value)
    rText.value = wanakana.toRomaji(kText.value)
})

document.getElementById("romaji").addEventListener("click", async () => {
    try {
        await navigator.clipboard.writeText(rText.value)

        clipInfo.innerHTML = "Copied romaji to clipboard!"
        setTimeout(() => { clipInfo.innerHTML = "" }, clipPopupTime)
      } catch (err) {
        console.error('Failed to copy romaji: ', err)
      }
})

document.getElementById("hiragana").addEventListener("click", async () => {
    try {
        await navigator.clipboard.writeText(hText.value)

        clipInfo.innerHTML = "Copied hiragana to clipboard!"
        setTimeout(() => { clipInfo.innerHTML = "" }, clipPopupTime)
      } catch (err) {
        console.error('Failed to copy hiragana: ', err)
      }
})


document.getElementById("katakana").addEventListener("click", async () => {
    try {
        await navigator.clipboard.writeText(kText.value)

        clipInfo.innerHTML = "Copied katakana to clipboard!"
        setTimeout(() => { clipInfo.innerHTML = "" }, clipPopupTime)
      } catch (err) {
        console.error('Failed to copy katakana: ', err)
      }
})


