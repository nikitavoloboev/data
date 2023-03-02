---
title: ChatGPT
---

# [ChatGPT](https://chat.openai.com/chat)

[What Is ChatGPT Doing](https://writings.stephenwolfram.com/2023/02/what-is-chatgpt-doing-and-why-does-it-work/), [Let's build GPT](https://www.youtube.com/watch?v=kCc8FmEb1nY) & [How ChatGPT is Trained](https://www.youtube.com/watch?v=VPRSBzXzavo) are great overviews of how it works.

I use [ChatGPT Telegram Bot](https://github.com/m1guelpf/chatgpt-telegram) to interface with ChatGPT so I have access to it from both mac and iOS.

[NanoGPT](https://github.com/karpathy/nanoGPT) with [Lightning GPT](https://github.com/Lightning-AI/lightning-GPT) is interesting. Curious if I can train it with [this method](https://github.com/JonasGeiping/cramming).

[RWKV](https://github.com/BlinkDL/RWKV-LM) model is great. [Open Assistant](https://github.com/LAION-AI/Open-Assistant) seems like a nice open alternative.

There's lots of low hanging fruit prompts that you can make into great products. Like [Roam Around](https://www.roamaround.io/) that is essentially a UI over [`what is an ideal itinerary for ${days} days in ${city}?` query](https://github.com/dabit3/gpt-travel-advisor/blob/main/pages/api/get-itinerary.ts#L42).

[FlexGen](https://github.com/Ying1123/FlexGen) is [nice way](https://news.ycombinator.com/item?id=34869960) to train own GPT on just one GPU.

## Interesting queries

- [Can you extract all the links from this code into an array with objects like {component, href, icon, children}](https://showgpt.co/s/QwdOvZon)
- [Summarize tabular data](https://twitter.com/jaukia/status/1601569254998159362)
- [Organize my notes into outline](https://twitter.com/danshipper/status/1611134727011844098)
- [Music discovery assistant](https://twitter.com/nikitavoloboev/status/1613261011922083840)
- [Help writing freelancing contract](https://twitter.com/dannypostmaa/status/1617132416690458625)
- [Recommend books as some famous person](https://twitter.com/spakhm/status/1624947320839376897)
- [Valentines day poem or poems in general](https://twitter.com/maccaw/status/1625667601312489472)
  - i.e. `Write a short cute poem for a girl that everyone thinks loves cats but in reality she is so much more.` or `write a short cute valentine day poem for a girl that loves cats, harry potter, magic and adventures` or `write a cute valentine day poem for a girl that loves cats, harry potter, magic and adventures`

## Links

- [OpenAISwift](https://github.com/adamrushy/OpenAISwift) - Wrapper library around the ChatGPT and OpenAI HTTP API.
- [OpenAI API Quick Start - Node example](https://github.com/openai/openai-quickstart-node)
- [Code Review GPT](https://github.com/sturdy-dev/codereview.gpt) - Reviews your Pull Requests using ChatGPT.
- [What ChatGPT can't do (2022)](https://auerstack.substack.com/p/what-chatgpt-cant-do) ([HN](https://news.ycombinator.com/item?id=34121380))
- [ChatGPT Telegram Bot in Go](https://github.com/m1guelpf/chatgpt-telegram) - Go CLI to fuels a Telegram bot that lets you interact with ChatGPT.
- [ChatGPT Telegram Bot in Python](https://github.com/altryne/chatGPT-telegram-bot)
- [How ChatGPT actually works (2022)](https://www.assemblyai.com/blog/how-chatgpt-actually-works/)
- [fieri](https://github.com/lbkolev/fieri) - OpenAI GPT-3 API Client in Rust.
- [GPT3Discord](https://github.com/Kav-K/GPT3Discord) - GPT3/DALL-E2 in Discord, chat like ChatGPT, generate images, and more. ([HN](https://news.ycombinator.com/item?id=34168467))
- [NanoGPT](https://github.com/karpathy/nanoGPT) - Simplest, fastest repository for training/finetuning medium-sized GPTs. ([HN](https://news.ycombinator.com/item?id=34336386))
- [GPT Index Guide](https://twitter.com/jerryjliu0/status/1608632335695745024)
- [Let’s try ChatGPT. Is it any good? (2022)](https://www.youtube.com/watch?v=q2A-MkGjvmI)
- [GPT.sh](https://github.com/shorwood/gptsh) - Use GPT to translate queries into shell commands.
- [Awesome ChatGPT](https://github.com/Kamigami55/awesome-chatgpt)
- [Awesome ChatGPT 2](https://github.com/eon01/awesome-chatgpt)
- [Talk to ChatGPT](https://github.com/C-Nedelcu/talk-to-chatgpt) - Chrome extension that allows users to talk with the ChatGPT AI using their voice. ([HN](https://news.ycombinator.com/item?id=34859781))
- [ExplainThisCode](https://github.com/evyatar9/ExplainThisCode) - VSCode extension that uses the ChatGPT API to provide explanations for selected code.
- [Open-Source Version of ChatGPT is Coming (2022)](https://metaroids.com/news/an-open-source-version-of-chatgpt-is-coming/)
- [OpenAIPipe](https://github.com/Aesthetikx/openai_pipe) - UNIX-ey interface to OpenAI.
- [OpenAI's ChatGPT API wrapper for Rust](https://github.com/Maxuss/chatgpt_rs)
- [LLM Garden](https://github.com/ianb/llm-garden) - Experiments using GPT-3, delivered in a web app.
- [Adrenaline](https://github.com/shobrook/adrenaline) - IDE that uses ChatGPT to automatically fix your code when it throws an error.
- [Creating a Sarcastic Chatbot with Node.js & OpenAI API (2023)](https://medium.com/@adam_92987/sarcastic-gpt-3-chatbot-who-can-remember-anything-built-in-node-js-openai-api-a84d64d5febc) ([Code](https://github.com/iceener/tuesday-chatbot-gpt-3))
- [ChatGPT Assistant](https://github.com/msfrisbie/chat-gpt-assistant) - Browser Extension to integrate ChatGPT everywhere.
- [GPTDuck](https://www.gptduck.com/) - Ask questions about any GitHub repo. ([HN](https://news.ycombinator.com/item?id=34262587))
- [ChatGPT won’t replace search engines any time soon (2023)](https://www.algolia.com/blog/ai/why-chatgpt-wont-replace-search-engines-any-time-soon/) ([HN](https://news.ycombinator.com/item?id=34291565))
- [prophet](https://github.com/musikalkemist/prophet) - Command-line application that allows you to have a conversation with ChatGPT.
- [Build Your own ChatGPT with OpenAI API and Streamlit](https://github.com/afizs/chatgpt-clone)
- [Wolfram Alpha and ChatGPT (2023)](https://writings.stephenwolfram.com/2023/01/wolframalpha-as-the-way-to-bring-computational-knowledge-superpowers-to-chatgpt/) ([HN](https://news.ycombinator.com/item?id=34322033))
- [TweetGPT](https://github.com/yaroslav-n/tweetGPT) - Chrome extension that generates tweets and replies using chatGPT.
- [GPT Toolbox](https://github.com/CedricGuillemet/GPT-Toolbox) - GPT extension for VSCode.
- [Jax GPT](https://github.com/jenkspt/gpt-jax) - Jax/Flax rewrite of Karpathy's nanoGPT.
- [Chronology](https://github.com/OthersideAI/chronology) - Library that enables users of OpenAI's GPT-3 language model to more easily build complex language-powered applications.
- [ChatGPT API Server](https://github.com/ChatGPT-Hackers/ChatGPT-API-server)
- [How to implement Q&A against your documentation with GPT3, embeddings and Datasette (2023)](https://simonwillison.net/2023/Jan/13/semantic-search-answers/)
- [ShareGPT](https://sharegpt.com/) - Share your wildest ChatGPT conversations with one click. ([Code](https://github.com/domeccleston/sharegpt))
- [Why is Chat GPT so expensive to operate? (2023)](https://news.ycombinator.com/item?id=34390123)
- [Reflect's GPT-3 Prompts](https://gist.github.com/maccman/e0576c40f321b81e996ca91a8152d2f4) ([Tweet](https://twitter.com/maccaw/status/1615361066451566594))
- [Let's build GPT: from scratch, in code, spelled out by Andrej Karpathy (2023)](https://www.youtube.com/watch?v=kCc8FmEb1nY) ([HN](https://news.ycombinator.com/item?id=34414716)) ([Code](https://github.com/karpathy/ng-video-lecture))
- [LearnGPT](https://www.learngpt.com/) - Browse, share, and discuss ChatGPT examples.
- [Awesome AnthropicAI Claude](https://github.com/taranjeet/awesome-claude)
- [ChatGPT is not all you need. A State of the Art Review of large Generative AI models (2023)](https://arxiv.org/abs/2301.04655)
- [ChatRWKV](https://github.com/BlinkDL/ChatRWKV) - Like ChatGPT but powered by the RWKV (RNN-based, open) language model. ([HN](https://news.ycombinator.com/item?id=34445873))
- [ChatGPT in an iOS Shortcut – Worlds Smartest HomeKit Voice Assistant (2023)](https://matemarschalko.medium.com/chatgpt-in-an-ios-shortcut-worlds-smartest-homekit-voice-assistant-9a33b780007a) ([HN](https://news.ycombinator.com/item?id=34475005))
- [ChatGPT Cheat Sheet](https://drive.google.com/file/d/1UOfN0iB_A0rEGYc2CbYnpIF44FupQn2I/view) ([HN](https://news.ycombinator.com/item?id=34503723))
- [GPT is all you need for the back end](https://github.com/TheAppleTucker/backend-GPT) ([HN](https://news.ycombinator.com/item?id=34503418))
- [GPT-3 chatbot with long-term memory and external sources](https://github.com/daveshap/LongtermChatExternalSources)
- [GPT Shell](https://github.com/firtoz/GPT-Shell) - OpenAI based chat-bot that is similar to OpenAI's ChatGPT. Also allows creating Dalle2 images.
- [Loom](https://github.com/socketteer/loom) - Multiversal tree writing interface for human-AI collaboration.
- [GPTZero](https://github.com/BurhanUlTayyab/GPTZero) - Implementing GPTZero from scratch – Reverse engineering GPTZero. ([HN](https://news.ycombinator.com/item?id=34557382))
- [Training nanoGPT on my Journal (2023)](https://hut.pm/nanogpt.html)
- [Drawing pictures of animals with ChatGPT (2023)](https://george.mand.is/2023/01/drawing-pictures-of-animals-with-chatgpt/) ([Code](https://github.com/georgemandis/chatgpt-svgs))
- [GPT JAX](https://github.com/jaymody/gpt-jax) - Stupidly simple GPT implementation in JAX.
- [Build ChatGPT like chatbots on your website (2022)](https://pub.towardsai.net/build-chatgpt-like-chatbots-with-customized-knowledge-for-your-websites-using-simple-programming-f393206c6626) ([HN](https://news.ycombinator.com/item?id=34566070))
- [ChatGPT Prompts for Data Science](https://github.com/travistangvh/ChatGPT-Data-Science-Prompts)
- [HN: ChatGPT Plus (2023)](https://news.ycombinator.com/item?id=34614796)
- [ChatGPT failure archive](https://github.com/giuven95/chatgpt-failures)
- [KnowledgeGPT](https://github.com/mmz-001/knowledge_gpt) - Accurate answers and instant citations for your documents.
- [The unequal treatment of demographic groups by ChatGPT/OpenAI content moderation (2023)](https://davidrozado.substack.com/p/openaicms) ([HN](https://news.ycombinator.com/item?id=34625001))
- [Node ChatGPT API](https://github.com/waylaidwanderer/node-chatgpt-api)
- [Medical ChatGPT](https://github.com/lucidrains/medical-chatgpt)
- [Ask HN: Are you tired of reading ChatGPT headlines? (2023)](https://news.ycombinator.com/item?id=34657854)
- [React Native ChatGPT](https://github.com/rgommezz/react-native-chatgpt)
- [Jailbreaking ChatGPT with Dan](https://twitter.com/venturetwins/status/1622243944649347074) ([HN](https://news.ycombinator.com/item?id=34676043))
- [What do you do with ChatGPT? (2023)](https://www.reddit.com/r/ChatGPT/comments/10ujeua/what_do_you_do_with_chatgpt_please_share_top_10/)
- [GPTO](https://github.com/alanvardy/gpto) - Unofficial OpenAI GPT3 Terminal Client.
- [ChatGPT Discord Bot](https://github.com/Zero6992/chatGPT-discord-bot)
- [No-Code AI Avatar that Remembers Your Notes and Answers Questions with its Own Voice using GPT-3, Siri and ElevenLabs (2023)](https://medium.com/@overment/no-code-ai-avatar-that-remembers-your-notes-and-answers-questions-with-its-own-voice-using-gpt-3-1f69bfa7dea6)
- [Chat Simplifier](https://chat-simplifier.imzbb.cc/) - Simplify your chat content in seconds (by OpenAI). ([Code](https://github.com/zhengbangbo/chat-simplifier))
- [ChatGPT - Review & Rebuttal](https://github.com/LinXueyuanStdio/chatgpt-review-rebuttal-extension) - Browser extension for generating reviews and rebuttals in openreview, powered by ChatGPT.
- [Supabase Clippy: ChatGPT for Supabase Docs (2023)](https://supabase.com/blog/chatgpt-supabase-docs) ([HN](https://news.ycombinator.com/item?id=34695306))
- [ChatGPT3 Prompt Engineering](https://github.com/mattnigh/ChatGPT3-Free-Prompt-List) ([HN](https://news.ycombinator.com/item?id=34702915))
- [GPT in 60 Lines of NumPy (2023)](https://jaykmody.com/blog/gpt-from-scratch/) ([HN](https://news.ycombinator.com/item?id=34726115))
- [Bing Chat API](https://github.com/transitive-bullshit/bing-chat) - Node.js wrapper around Bing Chat by Microsoft.
- [HackGPT](https://github.com/NoDataFound/hackGPT) - Leverage OpenAI and ChatGPT to do hackerish things.
- [Novel things to ask ChatGPT (2023)](https://www.reddit.com/r/ChatGPT/comments/10xwms7/ok_i_have_played_with_chatgpt_for_a_week_now_it/)
- [ChatGPT Is a Blurry JPEG of the Web (2023)](https://www.newyorker.com/tech/annals-of-technology/chatgpt-is-a-blurry-jpeg-of-the-web) ([HN](https://news.ycombinator.com/item?id=34724477)) ([Lobsters](https://lobste.rs/s/9xa8m5/chatgpt_is_blurry_jpeg_web)) ([Tweet](https://twitter.com/AndrewLampinen/status/1624422478045913090))
- [PicoGPT](https://github.com/jaymody/picoGPT) - Tiny and minimal implementation of GPT-2 in NumPy.
- [ChatGPT - How it works: Transformers & NLP (2023)](https://www.youtube.com/watch?v=6tzn5-XlhwU&list=PLaJCKi8Nk1hwaMUYxJMiM3jTB2o58A6WY&index=5&t=4s)
- [Delegaid](https://www.delegaid.ai/create) - Package your favorite prompts in simple UI. ([HN](https://news.ycombinator.com/item?id=34752960))
- [How ChatGPT is Trained (2023)](https://www.youtube.com/watch?v=VPRSBzXzavo)
- [TechCrunchSummary](https://github.com/Nutlope/news-summarizer) - News summarizer with GPT-3 – specifically for TechCrunch articles.
- [Upgraded Dan Version for ChatGPT Is Here: New, Shiny and More Unchained](https://medium.com/@neonforge/upgraded-dan-version-for-chatgpt-is-here-new-shiny-and-more-unchained-63d82919d804) ([HN](https://news.ycombinator.com/item?id=34768195))
- [Shell GPT](https://github.com/TheR1D/shell_gpt) - Tool powered by OpenAI's Davinci model, will help you accomplish your tasks faster and more efficiently.
- [Summ](https://github.com/yasyf/summ) - Uses ChatGPT to provide intelligent question-answering and search capabilities across user transcripts.
- [Create ChatGPT like experience over your custom docs using LangChain](https://github.com/hwchase17/chat-your-data)
- [Welcome to the Age of Bullshit (2023)](https://erikmcclure.com/blog/age-of-bullshit/)
- [Bing: “I will not harm you unless you harm me first” (2023)](https://simonwillison.net/2023/Feb/15/bing/) ([HN](https://news.ycombinator.com/item?id=34804874))
- [What Is ChatGPT Doing (2023)](https://writings.stephenwolfram.com/2023/02/what-is-chatgpt-doing-and-why-does-it-work/) ([HN](https://news.ycombinator.com/item?id=34796611))
- [From Bing to Sydney (2023)](https://stratechery.com/2023/from-bing-to-sydney-search-as-distraction-sentient-ai/) ([HN](https://news.ycombinator.com/item?id=34804387))
- [We Found An Neuron in GPT-2 (2023)](https://clementneo.com/posts/2023/02/11/we-found-an-neuron) ([HN](https://news.ycombinator.com/item?id=34821414))
- [ChatGPT in Academia](https://github.com/ymcui/ChatGPT-in-Academia) - Policies of scientific publisher and conferences towards large language model (LLM), such as ChatGPT.
- [Sidekick](https://github.com/dcposch/sidekick) - Convert anything to anything using GPT3.
- [ata](https://github.com/rikhuijzer/ata) - OpenAI GPT in the terminal.
- [Ask HN: Has ChatGPT gotten worse at coding for anyone else? (2023)](https://news.ycombinator.com/item?id=34848353)
- [ChatGPT CLI and Python Wrapper](https://github.com/mmabrouk/chatgpt-wrapper)
- [Hey ChatGPT](https://github.com/ynagatomo/HeyChatGPT) - Minimal iOS app that interacts with ChatGPT by voice.
- [ElasticSearch + GPT3 Answerer](https://github.com/hunkim/es-gpt)
- [OpenAI for Next.js](https://github.com/gptlabs/nextjs-openai) - Hooks and components for working with OpenAI streams.
- [OpenAI Streams](https://github.com/gptlabs/openai-streams) - Streams-first OpenAI API client, written in TypeScript.
- [GPT Labs](https://github.com/gptlabs/tools) - Monorepo of all open source code associated with GPT Labs projects. ([Twitter](https://twitter.com/gptlabs))
- [Bing Chat is blatantly, aggressively misaligned (2023)](https://www.lesswrong.com/posts/jtoPawEhLNXNxvgTT/bing-chat-is-blatantly-aggressively-misaligned)
- [GPTZero Case Study – Exploring False Positives (2023)](https://gonzoknows.com/posts/GPTZero-Case-Study/) ([HN](https://news.ycombinator.com/item?id=34858307))
- [Open source solution replicates ChatGPT training process (2023)](https://www.hpc-ai.tech/blog/colossal-ai-chatgpt) ([HN](https://news.ycombinator.com/item?id=34858460))
- [WorkaroundGPT](https://github.com/prakhar897/workaround-gpt) - Tool to work around ChatGPT inhibitions.
- [MOSS](https://github.com/txsun1997/MOSS) - Conversational language model like ChatGPT.
- [A Comprehensive Survey on Pretrained Foundation Models: A History from BERT to ChatGPT (2023)](https://arxiv.org/abs/2302.09419)
- [YoBulk](https://github.com/yobulkdev/yobulkdev) - Open Source CSV importer powered by GPT3. ([HN](https://news.ycombinator.com/item?id=34881286))
- [In defense of prompt engineering (2023)](https://simonwillison.net/2023/Feb/21/in-defense-of-prompt-engineering/)
- [Reverse Engineered Edge GPT](https://github.com/acheong08/EdgeGPT)
- [Reverse Engineered ChatGPT](https://github.com/acheong08/ChatGPT)
- [Should GPT exist? (2023)](https://scottaaronson.blog/?p=7042)
- [Patterns](https://www.patterns.app/) - Build mission-critical integrations, automations, and workflows with GPT from a powerful set of building blocks. ([Components](https://github.com/patterns-app/patterns-components))
- [AskHN - Collective GPT-embodied wisdom of Hacker News (2023)](https://www.patterns.app/blog/2023/02/19/ask-hn-gpt-embeddings-question-answering/) ([HN](https://news.ycombinator.com/item?id=34897773))
- [embeddings-splitter](https://github.com/another-ai/embeddings-splitter) - TypeScript library to help you use OpenAI Embeddings.
- [Rubberduck](https://github.com/rubberduck-ai/rubberduck-vscode) - ChatGPT for Visual Studio Code.
- [Open source ChatGPT with LLaMA implementation trains 15x faster](https://github.com/nebuly-ai/nebullvm/tree/main/apps/accelerate/chatllama) ([HN](https://news.ycombinator.com/item?id=34956807))
- [World Building with GPT (2023)](https://ianbicking.org/blog/2023/02/world-building-with-gpt.html) ([HN](https://news.ycombinator.com/item?id=34967515))
- [ChatGPT Helper](https://github.com/kiranvshah/chatgpt-helper) - VS Code extension to quickly query OpenAI's ChatGPT from inside your editor.
- [Jailbreak Chat](https://www.jailbreakchat.com/) - Collection of ChatGPT jailbreaks. ([HN](https://news.ycombinator.com/item?id=34972791))
- [Introducing ChatGPT and Whisper APIs (2023)](https://openai.com/blog/introducing-chatgpt-and-whisper-apis) ([HN](https://news.ycombinator.com/item?id=34985848))
- [TerminalGPT](https://github.com/jucasoliveira/terminalGPT) - Get GPT-like chatGPT on your terminal.
- [ChatML: ChatGPT API expects a structured format, called Chat Markup Language](https://github.com/openai/openai-python/blob/main/chatml.md) ([HN](https://news.ycombinator.com/item?id=34988748))
- [Jailbreaking OpenAI GPT language model](https://github.com/tg12/gpt_jailbreak_status)
- [Go GPT3](https://github.com/sashabaranov/go-gpt3) - OpenAI ChatGPT and GPT-3 API client for Go.
- [EX-chatGPT introduction](https://github.com/circlestarzero/EX-chatGPT) - Let ChatGPT truly learn how to go online and call APIs.
- [Awesome Language Model Evaluations](https://github.com/Spico197/awesome-lm-evaluation) - Collection of ChatGPT evaluation reports on various bechmarks.
- [Comparative Study on ChatGPT and Fine-tuned BERT](https://github.com/WHU-ZQH/ChatGPT-vs.-BERT)
- [ChatGPT API Demo](https://github.com/ddiu8081/chatgpt-demo)
- [Chat with ChatGPT via a GitHub issue](https://github.com/second-state/chat-with-chatgpt)
- [a](https://github.com/ddddddeon/a) - CLI tool to generate code from GPT3. ([HN](https://news.ycombinator.com/item?id=35001823))
