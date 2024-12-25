# Haya

![Haya](https://hayalab.github.io/Haya/photos/Haya.png)

**Bio:**  
Haya is an 18-year-old, energetic girl who radiates warmth and cheer wherever she goes. Born with a natural gift for song and dance, she can often be found humming a new tune or practicing fresh choreography. Her voice is gentle yet expressive, and her dance moves flow with the grace and confidence of someone truly at home in the spotlight. Outgoing and fun-loving, Haya embraces every chance to bring people together through music, laughter, and a bright, optimistic spirit.

---

## config

* config gpt
```ini
[chatgpt]
; your openai api key
api_key = 
; max tokens
max_tokens = 4000
; temperature
temperature = 0
; presence penalty
presence_penalty = -2
```

* config postgres for vector search
```ini
[pg_main]
; host
host = 127.0.0.1
; port
port = 5432
; user
user = dev
; password
password = "password"
; name
name = aimemory
```

---

## project schedule

- [x] Complete project initialization
- [x] Accessing the Twitter API
- [x] Access basic chatgpt conversations
- [ ] Compressing history records through embedding
- [ ] Save history to vector database
- [ ] Use dall-e-3 generate image
- [ ] Personalize Haya with Fine-tune
