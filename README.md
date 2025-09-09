flowchart TD

    %% API Layer
    API[API]

    %% Messaging
    Kafka[(Kafka)]
    Workers[Workers]

    %% Database
    Postgres[(Postgres)]

    %% Entities
    Election["
    election:
    - id
    - name
    "]

    Candidate["
    candidates:
    - id
    - name
    - election_id
    "]

    Vote["
    vote:
    - id
    - datetime
    - election_id
    - candidate_id
    "]

    %% API endpoints
    ElectionOps["create/get/update/delete election"]
    CandidateOps["create/get/update/delete candidate"]
    VoteOps["create vote / get all votes"]

    %% Relationships
    API --> Kafka
    API --> Postgres

    Kafka --> Workers
    Workers --> Postgres

    %% Domain relations
    Election --- Candidate
    Candidate --- Vote

    %% Attach operations to entities
    Election --> ElectionOps
    Candidate --> CandidateOps
    Vote --> VoteOps
