
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE user_role AS ENUM ('Admin', 'Employee', 'AssetManager', 'EmployeeManager');
CREATE TYPE employment_type AS ENUM ('Intern', 'Freelancer', 'FullTime');

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       fname TEXT NOT NULL,
                       lname TEXT,
                       email TEXT NOT NULL,
                       phone_no TEXT,
                       created_by UUID,
                       role user_role NOT NULL DEFAULT 'Employee',
                       emp_type employment_type NOT NULL DEFAULT 'FullTime',
                       created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMPTZ,
                       is_deleted BOOLEAN DEFAULT FALSE
);
CREATE UNIQUE INDEX idx_users_email_unique
    ON users (email)
    WHERE is_deleted = FALSE;


CREATE UNIQUE INDEX idx_users_phone_no_unique
    ON users (phone_no)
    WHERE is_deleted = FALSE AND phone_no IS NOT NULL;


--- table no.2

CREATE TYPE asset_type AS ENUM ('laptop', 'mouse', 'harddisk', 'pendrive', 'mobile', 'monitor', 'accessories');
CREATE TYPE asset_status AS ENUM ('assigned', 'available', 'waiting_for_service', 'in_service', 'damaged');



CREATE TABLE assets (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        brand TEXT NOT NULL,
                        model TEXT NOT NULL,
                        asset_type asset_type NOT NULL,
                        category TEXT,
                        owned_by TEXT NOT NULL DEFAULT 'Remotestate',
                        purchase_price NUMERIC(12,2),
                        purchased_date DATE NOT NULL,
                        warranty_start DATE,
                        warranty_expire DATE,
                        asset_status asset_status NOT NULL DEFAULT 'available',
                        is_active BOOLEAN DEFAULT TRUE,
                        is_retired BOOLEAN DEFAULT FALSE,
                        created_by UUID,
                        created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                        deleted_at TIMESTAMPTZ
);

CREATE TABLE laptop (
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              asset_id UUID UNIQUE REFERENCES assets(id) ON DELETE CASCADE,
                              processor TEXT NOT NULL,
                              ram TEXT NOT NULL,
                              storage TEXT NOT NULL,
                              os TEXT NOT NULL,
                              created_by UUID REFERENCES users(id),
                              created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mobile(
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              asset_id UUID UNIQUE REFERENCES assets(id) ON DELETE CASCADE,
                              imei TEXT UNIQUE NOT NULL,
                              ram TEXT NOT NULL,
                              storage TEXT NOT NULL,
                              created_by UUID REFERENCES users(id),
                              created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE monitor (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               asset_id UUID UNIQUE REFERENCES assets(id) ON DELETE CASCADE,
                               screen_size TEXT NOT NULL,
                               resolution TEXT NOT NULL,
                               panel_type TEXT NOT NULL,
                               created_by UUID REFERENCES users(id),
                               created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mouse(
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             asset_id UUID UNIQUE REFERENCES assets(id) ON DELETE CASCADE,
                             dpi TEXT NOT NULL,
                             connection_type TEXT NOT NULL,
                             created_by UUID REFERENCES users(id),
                             created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE harddisk(
                                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                asset_id UUID UNIQUE REFERENCES assets(id) ON DELETE CASCADE,
                                capacity TEXT NOT NULL,
                                disk_type TEXT NOT NULL,
                                created_by UUID REFERENCES users(id),
                                created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE pendrive (
                                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                asset_id UUID UNIQUE REFERENCES assets(id) ON DELETE CASCADE,
                                capacity TEXT NOT NULL,
                                created_by UUID REFERENCES users(id),
                                created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE accessories (
                                   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                   asset_id UUID UNIQUE REFERENCES assets(id) ON DELETE CASCADE,
                                   name TEXT NOT NULL,
                                   work TEXT,
                                   created_by UUID REFERENCES users(id),
                                   created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE TYPE asset_timeline_status AS ENUM ('assigned', 'retrieved');

CREATE TABLE asset_timeline (
                                id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                asset_id UUID REFERENCES assets(id) ON DELETE CASCADE,
                                assigned_to UUID REFERENCES users(id) ON DELETE SET NULL,
                                assigned_by UUID REFERENCES users(id) ON DELETE SET NULL,
                                assigned_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                returned_at TIMESTAMPTZ,
                                reason TEXT,
                                status asset_timeline_status NOT NULL DEFAULT 'assigned',
                                created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX idx_asset_timeline_assigned_to ON asset_timeline (assigned_to);

CREATE TABLE asset_history (
                               id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                               asset_id UUID REFERENCES assets(id) ON DELETE CASCADE,
                               old_status asset_status NOT NULL,
                               new_status asset_status NOT NULL,
                               employee_id UUID REFERENCES users(id) ON DELETE SET NULL,
                               performed_by UUID REFERENCES users(id) ON DELETE SET NULL,
                               performed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


