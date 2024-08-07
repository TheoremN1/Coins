using Microsoft.EntityFrameworkCore;
using DatabaseAPI.Database;
using DatabaseAPI.Database.Models;
using Microsoft.EntityFrameworkCore.Storage;
using Microsoft.EntityFrameworkCore.Infrastructure;
using Newtonsoft.Json;

namespace DatabaseAPI;

public class Program
{
	public static void Main(string[] args)
	{
		var builder = WebApplication.CreateBuilder(args);
		using (var streamReader = new StreamReader(Path.Join("Configs", "database.json")))
		{
			var json = streamReader.ReadToEnd();
			var items = JsonConvert.DeserializeObject<Dictionary<string, string>>(json)
				?? throw new NullReferenceException("���-�� �� ��� � ������ database.json");
			var stringConnection =
				$"Host={items["host"]};" +
				$"Port={items["port"]};" +
				$"Database={items["database"]};" +
				$"Username={items["username"]};" +
				$"Password={items["password"]}";
			builder.Services.AddSingleton(_ => {
				var options = new DbContextOptionsBuilder<DatabaseContext>();
				options.UseNpgsql(stringConnection);
				return new DatabaseContext(options.Options);
			});
		}
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
			foreach(var item in items)
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
        context.Users.Add(new User(){ 
			Name = "Ivan", Surname = "Pupkin", 
			Login = "rockstar13", Password = "qwerty", 
			RoleKey = "admin" });
        context.Achievements.Add(new Achievement() { 
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
		app.Run();
	}
}
