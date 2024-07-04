@echo on

REM Changer le répertoire de travail pour le dépôt Git
cd C:\Users\xplit\Desktop\Cours\Matières\B1\ForumGO

REM Vérifier si le répertoire de travail est correct
if not exist .git (
    echo Ce répertoire ne semble pas contenir de dépôt Git valide.
    pause
    exit /b
)

REM Changer de branche et effectuer un pull pour chaque branche
echo Passage à la branche Xplit et pull
git checkout Xplit
git pull

echo Passage à la branche Aylan et pull
git checkout Aylan
git pull

echo Passage à la branche Nico et pull
git checkout Nico
git pull

echo Passage à la branche Shems et pull
git checkout Shems
git pull

REM Retourner à la branche Xplit
echo Retour à la branche Xplit
git checkout Xplit

@pause