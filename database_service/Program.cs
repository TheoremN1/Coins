using Microsoft.EntityFrameworkCore;
using database_service.Database;
using database_service.Database.Models;
using Microsoft.EntityFrameworkCore.Storage;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Newtonsoft.Json;

var builder = WebApplication.CreateBuilder(args);

var host = Environment.GetEnvironmentVariable("POSTGRES_HOST");
var name = Environment.GetEnvironmentVariable("POSTGRES_NAME");
var port = Environment.GetEnvironmentVariable("POSTGRES_PORT");
var password = Environment.GetEnvironmentVariable("POSTGRES_PASSWORD");
var user = Environment.GetEnvironmentVariable("POSTGRES_USER");
var stringConnection =
    $"Host={host};" +
    $"Port={port};" +
    $"Database={name};" +
    $"Username={user};" +
    $"Password={password}";
Console.WriteLine($"StrCon: {stringConnection}");
builder.Services.AddSingleton(_ =>
{
    var options = new DbContextOptionsBuilder<DatabaseContext>();
    options.UseNpgsql(stringConnection);
    return new DatabaseContext(options.Options);
});
builder.Services.AddControllers();

var app = builder.Build();

var context = app.Services.GetRequiredService<DatabaseContext>();
var databaseCreator = (RelationalDatabaseCreator)context.Database.GetService<IDatabaseCreator>();
// TODO: \/ ��� ������ ������ \/
databaseCreator.EnsureDeleted();
databaseCreator.EnsureCreated();
//		 /\ ��� ������ ������ /\
using (var streamReader = new StreamReader(Path.Join("Configs", "roles.json")))
{
    var json = streamReader.ReadToEnd();
    var items = JsonConvert.DeserializeObject<List<Dictionary<string, string>>>(json)
        ?? throw new NullReferenceException("���-�� �� ��� � ������ roles.json");
    foreach (var item in items)
    {
        if (item is null)
            throw new NullReferenceException("���-�� �� ��� � ������ roles.json");
        context.Roles.Add(new Role() { Key = item["key"], Name = item["name"] });
    }
};
using (var streamReader = new StreamReader(Path.Join("Configs", "statuses.json")))
{
    var json = streamReader.ReadToEnd();
    var items = JsonConvert.DeserializeObject<List<Dictionary<string, string>>>(json)
        ?? throw new NullReferenceException("���-�� �� ��� � ������ statuses.json");
    foreach (var item in items)
    {
        if (item is null)
            throw new NullReferenceException("���-�� �� ��� � ������ statuses.json");
        context.Statuses.Add(new Status() { Key = item["key"], Name = item["name"] });
    }
};
// TODO: \/ ��� ������ ������ \/
// ��� ������ ����� ������ ��� ����� ������������
context.Users.Add(new User()
{
    Name = "Ivan",
    Surname = "Pupkin",
    Login = "rockstar13",
    Password = "qwerty",
    RoleKey = "admin"
});
context.Achievements.Add(new Achievement()
{
    Name = "��������� ������",
    Description = "�� ����� ������ ����������� � ���� ������? �����������! :)",
    Reward = 30
});
context.Merches.Add(new Merch()
{
    Name = "������",
    Description = "����: ������, �����: 30�",
    Price = 20
});
//		 /\ ��� ������ ������ /\
context.SaveChanges();

app.UseAuthorization();
app.MapControllers();

var servicePort = Environment.GetEnvironmentVariable("DATABASE_SERVICE_PORT");
app.Run($"http://[::]:{servicePort}");
