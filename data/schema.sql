-- Execute this first

-- Plant
CREATE TABLE Plant (
    PlantID SERIAL NOT NULL,
    Name VARCHAR(256) NOT NULL,
    TDS FLOAT NOT NULL,
    PH FLOAT NOT NULL,
    Lux FLOAT NOT NULL,
    LightsOnHour FLOAT NOT NULL,
    LightsOffHour FLOAT NOT NULL,
    PRIMARY KEY (PlantID)
);

-- Location
CREATE TABLE Location (
    LocationID SERIAL NOT NULL,
    City VARCHAR(256) NOT NULL,
    Province VARCHAR(256) NOT NULL,
    PRIMARY KEY (LocationID)
);

-- Users
CREATE TABLE User (
    UserID SERIAL NOT NULL,
    Username VARCHAR(256) NOT NULL,
    IsAdministrator BOOLEAN NOT NULL DEFAULT FALSE,
    Hash VARCHAR(256) NOT NULL,
    Salt VARCHAR(256) NOT NULL,
    CreatedAt TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (UserID)
);

-- Nutrient
CREATE TABLE Nutrient (
    NutrientID SERIAL NOT NULL PRIMARY KEY,
    Part INT NOT NULL,
    Nitrogen INT NOT NULL,
    Phosphorus INT NOT NULL,
    Potassium INT NOT NULL
);

-- PHUpUnit
CREATE TABLE PHUpUnit (
    PHUpUnitID SERIAL NOT NULL PRIMARY KEY
);

CREATE TABLE PHDownUnit (
    PHDownUnitID SERIAL NOT NULL PRIMARY KEY
);

-- ModuleGroup
CREATE TABLE ModuleGroup (
    ModuleGroupID SERIAL NOT NULL,
    ModuleGroupLabel VARCHAR(64) NOT NULL,
    PlantID INT NOT NULL,
    LocationID INT NOT NULL,
    Param_TDS FLOAT NOT NULL,
    Param_PH FLOAT NOT NULL,
    Param_Humidity FLOAT NOT NULL,
    OnAuto BOOLEAN NOT NULL,
    LightsOnHour FLOAT NOT NULL,
    LightsOffHour FLOAT NOT NULL,
    PRIMARY KEY (ModuleGroupID),
    FOREIGN KEY (PlantID) REFERENCES Plant (PlantID),
    FOREIGN KEY (LocationID) REFERENCES Location (LocationID)
);

-- SensorData
CREATE TABLE SensorData (
    ModuleID SERIAL NOT NULL,
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
    ModuleID SERIAL NOT NULL,
    ModuleGroupID INT NOT NULL DEFAULT 0,
    Token VARCHAR(256) NOT NULL,
    PRIMARY KEY (ModuleID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID)
);

-- NutrientUnit
CREATE TABLE NutrientUnit (
    NutrientUnitID SERIAL NOT NULL,
    ModuleID INT NOT NULL,
    ModuleGroupID INT NOT NULL,
    PHUpUnitID INT NOT NULL,
    PHDownUnitID INT NOT NULL,
    NutrientID INT NOT NULL,
    PRIMARY KEY (NutrientUnitID),
    FOREIGN KEY (PHUpUnitID) REFERENCES PHUpUnit (PHUpUnitID),
    FOREIGN KEY (PHDownUnitID) REFERENCES PHDownUnit (PHDownUnitID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID),
    FOREIGN KEY (ModuleID) REFERENCES Module (ModuleID),
    FOREIGN KEY (NutrientID) REFERENCES Nutrient (NutrientID)
);

-- GrowUnit
CREATE TABLE GrowUnit (
    GrowUnitID SERIAL NOT NULL,
    ModuleID INT NOT NULL,
    ModuleGroupID INT NOT NULL,
    Capacity INT,
    PRIMARY KEY (GrowUnitID),
    FOREIGN KEY (ModuleID) REFERENCES Module (ModuleID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID)
);

-- Device
CREATE TABLE Device (
    DeviceID SERIAL NOT NULL,
    IsOn BOOLEAN NOT NULL DEFAULT FALSE,
    ModuleID INT NOT NULL,
    Label VARCHAR(256) NOT NULL,
    PRIMARY KEY (DeviceID),
    FOREIGN KEY (ModuleID) REFERENCES Module (ModuleID)
);

-- SensorData_ModuleGroup
CREATE TABLE SensorData_ModuleGroup (
    ModuleGroupID SERIAL NOT NULL,
    Timestamp timestamp NOT NULL DEFAULT NOW(),
    Humidity FLOAT NOT NULL,
    Temperature FLOAT NOT NULL,
    PRIMARY KEY (ModuleGroupID, Timestamp),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID)
);

-- Permission
CREATE TABLE Permission (
    UserID INT NOT NULL,
    ModuleGroupID INT NOT NULL,
    PermissionLevel INT NOT NULL DEFAULT 0,
    PRIMARY KEY (UserID, ModuleGroupID),
    FOREIGN KEY (UserID) REFERENCES Users (UserID),
    FOREIGN KEY (ModuleGroupID) REFERENCES ModuleGroup (ModuleGroupID)
);

