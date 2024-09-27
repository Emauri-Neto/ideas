-- Database: ideas

-- DROP DATABASE IF EXISTS ideas;

CREATE DATABASE ideas
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'pt_BR.UTF-8'
    LC_CTYPE = 'pt_BR.UTF-8'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

-- Tabelas Independentes ou relacionamento many to one

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    expertise TEXT,
    interests TEXT[],
    roles TEXT[],
    status TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE time (
    id UUID PRIMARY KEY,
    date TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    week INTEGER NOT NULL,
    month INTEGER NOT NULL,
    year INTEGER NOT NULL
);

CREATE TABLE message (
    id UUID PRIMARY KEY,
    type TEXT NOT NULL,
    text TEXT NOT NULL,
    time_id UUID NOT NULL,
    CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE
);

CREATE TABLE private_space (
    id UUID PRIMARY KEY,
    subject TEXT NOT NULL,
    text TEXT NOT NULL,
    time_id UUID NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
); 

CREATE TABLE preference (
    id UUID PRIMARY KEY,
    description TEXT NOT NULL,
    time_id  UUID NOT NULL,
    CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE
);

-- Criar many to many com users
CREATE TABLE  study (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    objective TEXT NOT NULL,
    methodology TEXT NOT NULL,
    max_participants INTEGER NOT NULL,
    participation_type TEXT NOT NULL,
    time_id UUID NOT NULL,
    responsible_id UUID NOT NULL,
    CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE,
    CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Criar many to many com users
CREATE TABLE discussion_thread (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    max_participants INTEGER NOT NULL,
    discussion_deadline DATE NOT NULL,
    status TEXT NOT NULL,
    time_id UUID NOT NULL,
    responsible_id UUID NOT NULL,
    study_id UUID NOT NULL,
    CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE,
    CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE
);

-- Criar many to many com study
CREATE TABLE thematic_area (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
);

-- Criar many to many com participante
CREATE TABLE invitation (
	id UUID PRIMARY KEY,
	type TEXT NOT NULL,
	text TEXT NOT NULL,
	accept BOOLEAN NOT NULL,
	study_id UUID NOT NULL,
	thread_id UUID NOT NULL,
	time_id UUID NOT NULL,
	CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE,
	CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE,
	CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES discussion_thread (id) ON DELETE CASCADE
);

CREATE TABLE comment (
	id UUID PRIMARY KEY,
	description TEXT NOT NULL,
	parent_comment_id UUID REFERENCES comment(id) ON DELETE CASCADE
);


-- Criar many to many com tag
CREATE TABLE ideas (
	id UUID PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	qualifier TEXT NOT NULL
);

CREATE TABLE document (
	id UUID PRIMARY KEY,
	type TEXT NOT NULL,
	summary TEXT NOT NULL,
	keywords TEXT[],
	authors TEXT[] NOT NULL,
	text TEXT NOT NULL,
	access_level TEXT NOT NULL
);

CREATE TABLE tag (
	id UUID PRIMARY KEY, 
	name TEXT NOT NULL
);

CREATE TABLE line_suggestion (
	id UUID PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	accept BOOLEAN NOT NULL,
	study_id UUID NOT NULL,
	responsible_id UUID NOT NULL,
	time_id UUID NOT NULL,
	CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE,
	CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE,
	CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE fact_post (
	id UUID PRIMARY KEY,
	reported BOOLEAN NOT NULL,
	
	responsible_id UUID NOT NULL,
	study_id UUID NOT NULL,
	idea_id UUID,
	comment_id UUID,
	thread_id UUID NOT NULL,
	time_id UUID NOT NULL,

	CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE,
	CONSTRAINT fk_idea FOREIGN KEY (idea_id) REFERENCES ideas (id) ON DELETE CASCADE,
	CONSTRAINT fk_comment FOREIGN KEY (comment_id) REFERENCES comment (id) ON DELETE CASCADE,
	CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES discussion_thread (id) ON DELETE CASCADE, 
	CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE
);   

CREATE TABLE fact_report (
	id UUID PRIMARY KEY,
	
	responsible_id UUID NOT NULL,
	study_id UUID NOT NULL,
	thread_id UUID NOT NULL,
	document_id UUID,
	time_id UUID NOT NULL,

	CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE,
	CONSTRAINT fk_document FOREIGN KEY (document_id) REFERENCES document (id) ON DELETE CASCADE,
	CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES discussion_thread (id) ON DELETE CASCADE, 
	CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE
);

CREATE TABLE review (
	id UUID PRIMARY KEY,
	assessment BOOLEAN NOT NULL,
	comment TEXT NOT NULL,
	
	responsible_id UUID NOT NULL,
	document_id UUID,
	report_id UUID,
	time_id UUID NOT NULL,

	CONSTRAINT fk_responsible FOREIGN KEY (responsible_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_document FOREIGN KEY (document_id) REFERENCES document (id) ON DELETE CASCADE,
	CONSTRAINT fk_report FOREIGN KEY (report_id) REFERENCES fact_report (id) ON DELETE CASCADE,
	CONSTRAINT fk_time FOREIGN KEY (time_id) REFERENCES time (id) ON DELETE CASCADE
);

-- many to many

CREATE TABLE users_message (
	id SERIAL PRIMARY KEY,

	message_id UUID NOT NULL,
	sender_id UUID NOT NULL,
	receiver_id UUID NOT NULL,

	CONSTRAINT fk_sender FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_receiver FOREIGN KEY (receiver_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_message FOREIGN KEY (message_id) REFERENCES message (id) ON DELETE CASCADE
);

CREATE TABLE users_invitation (
	id UUID PRIMARY KEY,

	invitation_id UUID NOT NULL,
	sender_id UUID NOT NULL,
	receiver_id UUID NOT NULL,

	CONSTRAINT fk_sender FOREIGN KEY (sender_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_receiver FOREIGN KEY (receiver_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_invitation FOREIGN KEY (invitation_id) REFERENCES invitation (id) ON DELETE CASCADE
);

CREATE TABLE users_study (
	id UUID PRIMARY KEY,

	user_id UUID NOT NULL,
	study_id UUID NOT NULL,

	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE
);

CREATE TABLE users_thread (
	id UUID PRIMARY KEY,
	role TEXT NOT NULL,

	user_id UUID NOT NULL,
	thread_id UUID NOT NULL,

	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_thread FOREIGN KEY (thread_id) REFERENCES discussion_thread (id) ON DELETE CASCADE
);

CREATE TABLE users_preference (
	id UUID PRIMARY KEY,

	user_id UUID NOT NULL,
	preference_id UUID NOT NULL,

	CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
	CONSTRAINT fk_preference FOREIGN KEY (preference_id) REFERENCES preference (id) ON DELETE CASCADE      
);

CREATE TABLE tag_idea (
	id UUID PRIMARY KEY,

	tag_id UUID NOT NULL,
	idea_id UUID NOT NULL,

	CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE,
	CONSTRAINT fk_idea FOREIGN KEY (idea_id) REFERENCES ideas (id) ON DELETE CASCADE
);

CREATE TABLE tag_comment (
	id UUID PRIMARY KEY,

	tag_id UUID NOT NULL,
	comment_id UUID NOT NULL,

	CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE,
	CONSTRAINT fk_comment FOREIGN KEY (comment_id) REFERENCES comment (id) ON DELETE CASCADE
);

CREATE TABLE tag_document (
	id UUID PRIMARY KEY,

	tag_id UUID NOT NULL,
	document_id UUID NOT NULL,

	CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tag (id) ON DELETE CASCADE,
	CONSTRAINT fk_document FOREIGN KEY (document_id) REFERENCES document (id) ON DELETE CASCADE
);

CREATE TABLE document_line_suggestion (
	id UUID PRIMARY KEY,

	document_id UUID NOT NULL,
	line_suggestion_id UUID NOT NULL,
	
	CONSTRAINT fk_document FOREIGN KEY (document_id) REFERENCES document (id) ON DELETE CASCADE,
	CONSTRAINT fk_line_suggestion FOREIGN KEY (line_suggestion_id) REFERENCES line_suggestion (id) ON DELETE CASCADE
);

CREATE TABLE document_private_space (
	id UUID PRIMARY KEY,

	document_id UUID NOT NULL,
	private_space_id UUID NOT NULL,
	
	CONSTRAINT fk_document FOREIGN KEY (document_id) REFERENCES document (id) ON DELETE CASCADE,
	CONSTRAINT fk_private_space FOREIGN KEY (private_space_id) REFERENCES private_space (id) ON DELETE CASCADE
);

CREATE TABLE document_fact_post (
	id UUID PRIMARY KEY,

	document_id UUID NOT NULL,
	fact_post_id UUID NOT NULL,
	
	CONSTRAINT fk_document FOREIGN KEY (document_id) REFERENCES document (id) ON DELETE CASCADE,
	CONSTRAINT fk_fact_post FOREIGN KEY (fact_post_id) REFERENCES fact_post (id) ON DELETE CASCADE
);

CREATE TABLE thematic_area_study (
	id UUID PRIMARY KEY,

	study_id UUID NOT NULL,
	thematic_area_id UUID NOT NULL,
	
	CONSTRAINT fk_study FOREIGN KEY (study_id) REFERENCES study (id) ON DELETE CASCADE,
	CONSTRAINT fk_thematic_area FOREIGN KEY (thematic_area_id) REFERENCES thematic_area (id) ON DELETE CASCADE
);