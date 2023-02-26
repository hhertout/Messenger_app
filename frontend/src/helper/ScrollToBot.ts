export const scrollToBot = (id:string) => {
    const el = document.getElementById(id)!
    window.scrollTo(0, el.scrollHeight)
}