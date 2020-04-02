-- Execute this first

-- Plant
CREATE TABLE Plant (
    PlantID VARCHAR(256) NOT NULL,
    Name VARCHAR(256) NOT NULL,
    TDS FLOAT NOT NULL,
    PH FLOAT NOT NULL,
    Lux FLOAT NOT NULL,
    PRIMARY KEY (PlantID)
);

-- Location
CREATE TABLE Location (
    LocationID VARCHAR(256) NOT NULL,
    City VARCHAR(256) NOT NULL,
    Province VARCHAR(256) NOT NULL,
    PRIMARY KEY (LocationID)
);

-- Users
CREATE TABLE Users (
    UserID VARCHAR(256) NOT NULL,
    Username VARCHAR(256) NOT NULL,
    IsAdministrator BOOLEAN NOT NULL DEFAULT FALSE,
    Hash VARCHAR(256) NOT NULL,
    Salt VARCHAR(256) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (UserID)
);

-- Nutrient
CREATE TABLE Nutrient (
    NutrientID VARCHAR(256) NOT NULL PRIMARY KEY,
    Part INT NOT NULL,
    Nitrogen INT NOT NULL,
    Phosphorus INT NOT NULL,
    Potassium INT NOT NULL
);

-- PHUpUnit
CREATE TABLE PHUpUnit (
    PHUpUnitID VARCHAR(256) NOT NULL PRIMARY KEY
);

CREATE TABLE PHDownUnit (
    PHDownUnitID VARCHAR(256) NOT NULL PRIMARY KEY
);

-- ModuleGroup
CREATE TABLE ModuleGroup (
    ModuleGroupID VARCHAR(256) NOT NULL,
    PlantID VARCHAR(256) NOT NULL,
    LocationID VARCHAR(256) NOT NULL,
    Param_TDS FLOAT NOT NULL,
    Param_PH FLOAT NOT NULL,
    Param_Humidity FLOAT NOT NULL,
    OnAuto BOOLEAN NOT NULL,
    LightOnTime time NOT NULL,
    LightOffTime time NOT NULL,
    PRIMARY KEY (ModuleGroupID),
    FOREIGN KEY (PlantID) REFERENCES Plant (PlantID),
    FOREIGN KEY (LocationID) REFERENCES Location (LocationID)
);

-- SensorData
CREATE TABLE SensorData (
    ModuleID VARCHAR(256) NOT NULL,
    Timestamp TIMESTAMP NOT NULL DEFAULT NOW(),
    TDS FLOAT NOT NULL,
    PH FLOAT NOT NULL,
    SolutionTemperature FLOAT NOT NULL,
    ArrGrowUnitLux FLOAT ARRAY NOT NULL,
    ArrGrowUnitHumidity FLOAT ARRAY NOT NULL,
    ArrGrowUnitTemperature FLOAT ARRAY NOT NULL,
    PRIMARY KEY (Timestamp, ModuleID),
    FOREIGN KEY (ModuleID) REFERENCES Module (ModuleID)
);

-- Module
CREATE TABLE Module (
    ModuleID VARCHAR(256) NOT NULL,
    ModuleGroupID VARCHAR(256) NOT NULL,
    Token VARCHAR(256) NOT NULL,
    PRIMARY KEY (ModuleID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID)
);

-- NutrientUnit
CREATE TABLE NutrientUnit (
    NutrientUnitID VARCHAR(256) NOT NULL,
    ModuleID VARCHAR(256) NOT NULL,
    ModuleGroupID VARCHAR(256) NOT NULL,
    PHUpUnitID VARCHAR(256) NOT NULL,
    PHDownUnitID VARCHAR(256) NOT NULL,
    NutrientID VARCHAR(256) NOT NULL,
    PRIMARY KEY (NutrientUnitID),
    FOREIGN KEY (PHUpUnitID) REFERENCES PHUpUnit (PHUpUnitID),
    FOREIGN KEY (PHDownUnitID) REFERENCES PHDownUnit (PHDownUnitID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID),
    FOREIGN KEY (ModuleID) REFERENCES Modules (ModuleID),
    FOREIGN KEY (NutrientID) REFERENCES Nutrient (NutrientID)
);

-- GrowUnit
CREATE TABLE GrowUnit (
    GrowUnitID VARCHAR(256) NOT NULL,
    ModuleID VARCHAR(256) NOT NULL,
    capacity INT,
    PRIMARY KEY (GrowUnitID),
    FOREIGN KEY (ModuleID) REFERENCES Modules(ModuleID)
);

-- Device
CREATE TABLE Device (
    DeviceID VARCHAR(256) NOT NULL,
    StatusID VARCHAR(256) NOT NULL,
    ModuleID VARCHAR(256) NOT NULL,
    name VARCHAR(256) NOT NULL,
    PRIMARY KEY (DeviceID),
    FOREIGN KEY (ModuleID) REFERENCES Modules(ModuleID)
);

-- SensorData_ModuleGroup
CREATE TABLE SensorData_ModuleGroup (
    ModuleGroupID VARCHAR(256) NOT NULL,
    Timestamp timestamp NOT NULL,
    Humidity FLOAT NOT NULL,
    Temperature FLOAT NOT NULL,
    PRIMARY KEY (ModuleGroupID, Timestamp),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup(ModuleGroupID)
);

-- Permission
CREATE TABLE Permission (
    UserID VARCHAR(256) NOT NULL,
    ModuleGroupID VARCHAR(256) NOT NULL,
    Supervisor BOOLEAN NOT NULL,
    Monitor BOOLEAN NOT NULL,
    Control BOOLEAN NOT NULL,
    PRIMARY KEY (UserID, ModuleGroupID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup(ModuleGroupID)
);

