import { surpriseMePrompts } from '../constants'

export function getRandomPrompt(prompt){
    const randomText = Math.floor(Math.random() * surpriseMePrompts.length)
    const randomPrompt = surpriseMePrompts[randomText];
    if(randomPrompt===prompt) return getRandomPrompt(prompt);
    return randomPrompt;
}