# WasaText

WasaText è un'applicazione di messaggistica istantanea full-stack sviluppata come progetto universitario per il corso *Web and Software Architecture* (WASA) presso la Sapienza — Università di Roma. Il progetto implementa una piattaforma di chat in tempo reale ispirata a WhatsApp, con supporto per conversazioni dirette e di gruppo, reazioni ai messaggi e gestione completa dei profili utente.

---

## Architettura

Il sistema segue un'architettura client-server con separazione netta tra presentation layer e business logic. La comunicazione avviene tramite un'API RESTful documentata secondo lo standard OpenAPI 3.0.

```
┌─────────────────┐        HTTP/JSON        ┌─────────────────┐
│   Vue 3 + Vite  │ ──────────────────────► │   Go REST API   │
│   (Nginx :8080) │ ◄────────────────────── │   (port :3000)  │
└─────────────────┘                         └────────┬────────┘
                                                     │ SQL
                                            ┌────────▼────────┐
                                            │  SQLite (file)  │
                                            └─────────────────┘
```

### Backend

Il backend è scritto in **Go** e strutturato intorno al pattern *handler → service → repository*. Il routing HTTP è affidato a `julienschmidt/httprouter`, scelto per la sua leggerezza e le prestazioni in scenari con molte route parametriche. La persistenza è gestita da **SQLite 3** tramite il driver `mattn/go-sqlite3`, con foreign key enforcement abilitato esplicitamente. Il logging è strutturato con **logrus**.

L'autenticazione si basa su un Bearer Token semplificato: al momento del primo accesso il server genera un identificatore univoco per l'utente che viene poi trasmesso come header `Authorization: Bearer <id>` in ogni richiesta successiva. Non esiste una sessione server-side: il token è lo user ID stesso, validato dal middleware `bearerAuth` prima di ogni handler protetto.

Il database viene inizializzato con uno schema embedded nel binario — le tabelle vengono create al primo avvio tramite migrazioni idempotenti. Il file SQLite è configurabile via flag CLI, variabile d'ambiente o file YAML.

### Frontend

Il frontend è una **Single Page Application** costruita con **Vue 3** (Composition API) e bundlata con **Vite**. Il routing client-side è gestito da Vue Router 4. Le chiamate HTTP verso il backend usano **Axios** con un'istanza configurata centralmente in `src/services/axios.js`, che inietta automaticamente il token di autenticazione da localStorage.

L'interfaccia è organizzata in cinque view principali (Login, Home, Profile, Nuovo Gruppo, Lista Utenti) e una serie di componenti riusabili che coprono le funzionalità core: area chat, rendering dei messaggi, gestione dei membri, inoltro e reazioni.

In produzione il frontend è servito da **Nginx**, configurato per gestire il fallback delle route SPA e il proxy verso il backend.

---

## Funzionalità

| Area | Capacità |
|---|---|
| **Utenti** | Registrazione implicita al primo login, cambio username, foto profilo |
| **Chat dirette** | Apertura conversazione con qualsiasi utente, eliminazione chat |
| **Gruppi** | Creazione, aggiunta/rimozione membri, cambio nome e foto |
| **Messaggi** | Invio testo e immagini, eliminazione, inoltro verso altre chat |
| **Reazioni** | Aggiunta e rimozione di emoji su qualsiasi messaggio |
| **Stato lettura** | Marcatura automatica dei messaggi come letti/non letti |

---

## Stack tecnologico

| Layer | Tecnologia | Versione |
|---|---|---|
| Backend language | Go | 1.17+ |
| HTTP router | julienschmidt/httprouter | 1.3 |
| Database | SQLite 3 | — |
| Logging | sirupsen/logrus | 1.9 |
| Frontend framework | Vue | 3.2 |
| Build tool | Vite | 3 |
| HTTP client | Axios | 0.28 |
| Container runtime | Docker + Compose | — |
| Web server (prod) | Nginx | latest |
| Linter (Go) | golangci-lint | — |

---

## Struttura del repository

```
Wasa-Project-2024/
├── cmd/webapi/                  # Entrypoint del server Go
│   ├── main.go                  # Bootstrap: config, DB, routing, graceful shutdown
│   ├── load-configuration.go    # Parsing flag / env / YAML
│   └── cors.go                  # CORS middleware
│
├── service/
│   ├── api/                     # Handler HTTP (~25 endpoint)
│   ├── database/                # Data access layer (~28 file)
│   ├── globaltime/              # Astrazione del clock (testabilità)
│   └── utilitytool/             # Utility condivise
│
├── webui/
│   └── src/
│       ├── views/               # Pagine applicazione
│       ├── components/          # Componenti riusabili
│       ├── router/              # Definizione route Vue
│       └── services/            # Configurazione Axios
│
├── doc/
│   └── api.yaml                 # Specifica OpenAPI 3.0
│
├── Dockerfile.backend
├── Dockerfile.frontend
└── docker-compose.yml
```

---

## Deployment

Il progetto include configurazioni Docker pronte per l'uso. Il build del backend usa un'immagine multi-stage (`golang:1.19` → `debian:bookworm-slim`) per minimizzare la dimensione dell'immagine finale. Il build del frontend compila i sorgenti Vue con Vite e li copia in un'immagine Nginx.

`docker-compose.yml` orchestra i due servizi esponendo il backend sulla porta **3000** e il frontend sulla porta **8080**.

---

## API

La specifica completa dell'API è disponibile in [doc/api.yaml](doc/api.yaml) nel formato OpenAPI 3.0. Comprende gli endpoint per autenticazione, gestione utenti, conversazioni, messaggi, reazioni e gruppi — in totale circa 20 route.

---

## Autore

**Marco Linardi** — Sapienza Università di Roma
